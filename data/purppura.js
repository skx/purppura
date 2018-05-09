function safer(text)
{
  return text
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;")
      .replace(/'/g, "&#039;");
}

function update_alerts()
{
    $.getJSON( "/events" , function( data ) {
        var h = {};
        h['raised'] = 0
        h['pending'] = 0
        h['acknowledged'] = 0

        // Clear our table of past events - by removing all rows of
        // the table, except for the first (which is the header-row).
        for (var key in h) {
            $("#" + key + "_alerts").find("tr:gt(0)").remove();
        }

        $.each( data, function( key, val ) {

            // Bump the count of this type of alerts, for the title.
            var status = val['Status'];
            h[status] = h[status]+1

            // This is horrid - we want to show either "will raise at",
            // "last notified at", or nothing depending on the type.
            if ( status == "raised" ) {

                var d = Math.round( val['NotifiedAt'] * 1000)
                d = parseFloat(d);
                d = new Date( d  );

                // id | source | subject | last notified | action
                var t = "<tr class=\"click\"><td>" + val['ID'] + "</td>";
                t += "<td>" + val['Source'] + "</td>";
                t += "<td>"  + safer(val['Subject']) + "</td>";
                t += "<td>" + d  + "</td>";
                t += "<td><a href=\"/acknowledge/" + (val['ID']) + "\">ack</a>"
                t += " <a href=\"/clear/" + (val['ID']) + "\">clear</a>"
                t += "</td>"

                // Add the alert
                $("#" + val['Status'] + "_alerts").find('tbody').append(t)

                // Add the details.
                $("#" + val['Status'] + "_alerts").find('tbody')
                    .append("<tr style=\"display:none;\"><td></td><td colspan=\"4\"><p>" + val['Detail'] + "</p></td></tr>")
            }
            if ( status == "acknowledged" ) {

                // id | source | subject | actions
                var t = "<tr class=\"click\"><td>" + val['ID'] + "</td>";
                t += "<td>" + val['Source'] + "</td>";
                t += "<td>" + safer(val['Subject']) + "</td>";
                t += "<td><a href=\"/raise/" + (val['ID']) + "\">raise</a> <a href=\"/clear/" + (val['ID']) + "\">clear</a></td>"

                // Add the alert
                $("#" + val['Status'] + "_alerts").find('tbody').append(t)

                // Add the details.
                $("#" + val['Status'] + "_alerts").find('tbody')
                    .append("<tr style=\"display:none;\"><td></td><td colspan=\"4\"><p>" + val['Detail'] + "</p></td></tr>")

            }
            if ( status == "pending" ) {

                var d = Math.round( val['RaiseAt'] * 1000)
                d = parseFloat(d);
                d = new Date( d  );

                // id | source | subject | last notified | action
                var t = "<tr class=\"click\"><td>" + val['ID'] + "</td>";
                t += "<td>" + val['Source'] + "</td>";
                t += "<td>" + safer(val['Subject']) + "</td>";
                t += "<td>" + d  + "</td>";
                t += "<td><a href=\"/clear/" + (val['ID']) + "\">clear</a></td>"

                // Add the alert
                $("#" + val['Status'] + "_alerts").find('tbody').append(t)

                // Add the details.
                $("#" + val['Status'] + "_alerts").find('tbody')
                    .append("<tr style=\"display:none;\"><td></td><td colspan=\"4\"><p>" + val['Detail'] + "</p></td></tr>")
            }


            // Set the title.
            document.title = "Alerts [" + h['raised'] + "/" + h['acknowledged'] + "/" + h['pending'] + "]";

        });
    });
}
