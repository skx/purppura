//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a simple array containing one entry for each embedded
// resource.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

	var tmp EmbeddedResource

	tmp.Filename = "data/index.html"
	tmp.Contents = "1f8b08000000000004ffcc577b6fdb3812ff3f9f62ca14671ba924a7c9f592543290cb76d17772dbf4d074b13850e2d8624c930a49c97117fdee078a922c25d9bb6e6f7158f90ff3319cc76f5e64fce887f3b3cbab8b1790db9598edc4ee0f04958b84a024b31d803847cadc0020b6dc0a9c9d0ad4d6c0cfd3681a4d7f8923bfba038e6285964296536dd026a4b4f3e0a8e602e0b7245d61422a8eeb42694b2053d2a2b409597366f38461c5330cd66ef204b8e4965311988c0a4cf6078c726b8b006f4a5e25e4cc33092e3705f6585abcb59133e879a7d1c7cb1fb70a092e97a05124c4d88d4093235a02b9c679421c7b731245c6d26c59509b87a952d6584d8b8cc93053aba85b880ec2c3701a65c66cd7c2159761660c012e2d2e34b79b84989c1e1c1d0617abab63f1f4e86a717ef87afdee4d7a592de887bf7d7afff9f5c1f4dd9be31f4f3f5f5f7ff964c5cde6f3d9cddfaf9e7d7a967e7ab57cb7b7f76ab9e1f2fd1e814c2b6394e60b2e1342a5929b952a0d99791f984cf3c282d15942a228530cc3eb9b12f5a656db0f83fd70ff697858ab796dc82c8efca9c68d7ee2597c0714d77791b87e10882a7f2dbfec9ffff3157bbbb41b7efbf2eaf0e3f2e0fce58b35fbc7cd4551dc5cfd74b4f7d7d5b5a1e58bcbc5dbf3ecc546fe70ac2e5ee6f997fda3f76b029956c628cdef02f1b039765360e2c3e29a56d41b49bc955151eaa2283575bade3befe2a37fdc79d6430575e6c0af35f40029cd960bad4ac9824c09a54f76e7c7eef7bcd9ffdafc87c66285c15c298bba3b5d50c6b85c9cc0e1b4b885697bc87382dde3e38ecf3d39704750aa34431d58559cc0d3e2168c129cc12e1ebb5fcbf8616d4249abc0657d4fb3468554d06cd99e9e2b698335f2456e4f205582b51b9e6d1cd5b07998da887ad8059e0660cd2553eb5049a1288304e6a5cc2c57723c69f1e573180b9551b71ae6d4e4f0284960349ac0e3f188feec9377047b3024da8311f96534092d4dc72393abf568d22a5b9f63d4d2c0aac542b800a1694d9c72c9c6a34cf06c397ab2d505279dbf864212783cc6d052bd403b09a9b57a3c720a8d26a1295363f578bf93fa75f27ca7098528825a06970b50126c8ea0d51ad65c08f02ad56b0c2de5c284cda9c7632096a602094c42860217d4a25bd361cd8d3c0120eda8c311b04269616b81fb2aaac16b5d9be0691a3ba0d3d87d0e7d78d41087dc8c815002933bfc001e8f6dcecd24cc843268ec7864f568124abc6d87deaef180b98f1ab803cdc782518ba04a0dd4f71dac506fe0600a0633255987485953fecb536d591bb4afa4455d5131ee7098fc7a8f1cbe3e8183e9743aed94f20af54b491ce54d338c53c5363e6c63492bc804352621925629d5e0ff0286735a0adb8577cc7847e91a1fe51275301725671d0d0ca81a464e2aea1e0d409c96d6ba70a96b9a9f90a11a4d40bbdc15b430c8080ce23c537ebd5dae033721bb5e2801aa390df0b6a092214bc89c0a83cdaad35e2bd1593c500d20360595ad3246074a8a0d995dd6f9e5c0e18b3a81e3c8d17914db6f7094674a0629d575412ee8ff81348e3c947da562da9ae29109524d256b6f0a11995d34ad238ee8f65c1c315ef5a6cef79cb52172d755ad2f3a670d108d4bd153c12198521d485a0da80062c1673105577412a7d84bb542a7541c09be5505208e4a3198ff16ff56947645febe30e849136aa14a4b666febffff2e75804f6f12479276b8d5a835b9e562aece989e1e71feb4b98ec651fe746bd27d83024b53436a0fac369734ed31a9716ba5d0ccf20a89336c902cae2934c0ee6aca5d36cd7eaaff1f34d53be23739d06c29d55a205b2023b3d3deec7bb81528ddb581cc2efce03e0feff0a6750c4b8ca5699dce28870eaea17701db58db02e4e80b2a11e6942170092d621d77f7c5c5ec2f3235c5f3382aeeec58d7b57a9c9b8add1720106aaac058cd0b64cd2c5715ba56950a0cfcfda6db729d00a5f30a0cc5b9578b9ec5369fedc691cdebd10755ea0cb7d332bdc6cc76f3b7d45878af2c9f736470badd38ad2f23c61346560f25c551ad737fb117d6feeb401d04c083d00e0214be0dd03ed73f1dac7f207a6dc0ff61c0350cff7498d515062eb92be34df0feaf300e62b29934a93b48daef99d47ce2e651e36a871f767ef253e8bf7d0868e52edd92b6b7822ef0eb64f94fc5dfb5d46e5f042b163ceb0ebb6fd805065b839abf7def90a636c0996278b76bd6675c63f0fdb57d172fb8cdcbb47e5b9be56df78e24b3d38a72e18a95bbd67baafb95f96e336e3ce2d477dfefb411becdc833556cea96fe8d26d60e0b97b82ac2398fc8ec837bbcc21b5c15bfd3a09e7571e4a3c1054d1ca58a6d663b7194db9598edfc3b0000ffff2a89bb5697120000"
	tmp.Length = 4759
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/login.html"
	tmp.Contents = "1f8b08000000000004ffc457dd73dbb8117ff75fb1456e7af6a414fd95d49641ce5c7db9c9b7dd8baf13e70d2420121608c000284bb9b9ffbd0380a449596dd33cb4d08c08808bc5ee6f3f89fff4f3d5e5cdedf52ba85d23f23dec1f2088ac32c424caf70070cd08f51300ecb8132c7faf2a2ee15776df72c3284ee3ee1e788a86390292342c432bce1eb4320e41a9a463d265e88153576794ad78c992b0f80b70c91d2722b125112c3b0a77024446b5733af1f7ac3274199924371bcd462c1d5bbbd44b7d01654d8c652efbede697e4ace723b85c82612243d66d04b335630e416dd822439ebd9da7a975a45c6ae2ea59a194b3ce105d52392b55930e1be9c9ec74769896d63eeecd1a2e67a5b508b874ac32dc6d32646b7272769a5c37b7e7e2f8ecb6ba3a7dfbf0e15d71b3aac8a7bf7efef8e5edc9e18777e7bffcf4e5eeeeeb6727ee375f2eefff76fbf2f3cbe2f39be587e7cfdf2c375c7e7c8ea034ca5a6578c56586885472d3a8d6a23c026d4bc3b5036bca0ca569a9289bddddb7cc6c82d8719a1ccd8e8e67a741cc3b8b729cc6539dade222b2f80e28eeb691b8db09c4aa7e2bbf1e5dfde30d7dbf741bbe7e7d7bfadbf2e4eaf5ab07faf7fb6badef6f7f3d7bfea2b9b3a47d7553bdbf2a5f6de4cfe7eafa755d7f3d3afbf880a034ca5a65f83610dbea78fb82db689645aff09689aa42706ff83d40075090725919d54a9a944a28337fb638f7bf8beefd1fdd73661d5bb164a194636638ad09a55c5673383dd46b38ec0f454ef0ecfc7ce0f3e41ed8baa85086329338a5e770acd76095e0149eb173ffeb19ef966626c92af1a13992ac13a110a45cf6a7174abae481f1aa76732894a0fd8bc816a701b60853ef11230cefc88a443fe9a0fc611fa82adb8649070733c308ddec2f5a593aaee4fe01fc0e3fecffc8a56edddcb1b59bafb8e58560f30537d6fd78305ba8b2b5fb0717f0c74100796c439c7a7dbc28b85074e3277e604956500a626d86245915c4407c24942d482b7ac9fcc0940fb43eeb102e994916a2e5b4933f8e315dc7ccdfcdcc840a0017ad734a7688c4059a0a933855558279f005d19651049438d26d67a85471bfdf26a6622e43cfe2b50888e124616b4d246534430b222ceb76bd06468941ef2de100b0d544f6e2589328293628bf890249b2e215f176c1a9a7ebf1ecc7e4302f954c0a624292d0e47f468cd308ea54384c7aa5224a496188a47dda4e517edd1aad5b43704ac627714af96ab2e1fd81d3de71b64dd7db6630de16c2b81523413ca205318924ab2d3a002c788e09c4ba92a2fcb56a98170ea7828f0502c0692bc63b5b324f9638956495ef85488189d37ad708ce3d92645056f8025da8358210dc196a88a9b80c79e6c5a15e5f0c3834847b42af7fd2d0e4653f518b85652e39096bdb2467fda47b713cba762a9826920908ff09970b8560acec0eda10785c56138e7e0475bab88ffc42a38162ff31c1298e5d5bdefe1d085dde0e289c1ceaf5004264ee53ceb6b09db81d03caad166433974a3214bc2a009d10c142971353545841f84f289115333d7647c73ebab63cd40fbc50a67964e85783707e91d4caf0af4a3a221018255816b611909076339406411034ccd58a66e8faead3cd133ca7ca742e5128e7543387e3172340b8cfdf89af907a0797adb433224e08a54aa21cf3de6e95d8e8dae7161866496b7d9ac529cf77e7253f70e01a20f1cd24ea12b0636bdf4e46a003323e0e8c12a8eb3923f18a889665088116a464b51294990cddaad680bfdc13c10ec576f88f1f631ffaffa12654b9f4eef38da86962ed8332b447ee71fd6fd07b247a0adcf06e1ace7e7c3b6ebed1391a875eb0617034d0ad1089f17dca8e2b3adfed642f95f049e9e8387c59f80ab9f308f46e14fb19db160d773170fbb977856c58757e13be6f86002c9c84c2c998cc9e2aef13fa8e90fe17db38f51a4fb93c816f6b63b2ec167d45d0f99f6561f5054e75fe3d8bc007770daeaf93713a281f9730ee83fb0cf4d85e0ca88448196c247717a8e17dac38c3613fa6f576f22a94d8fef063ef8bf24faa3525834b45d98e623b2ecbfd374ec55ddd16e13bc92ed7a9ee7a0994ffb4225c904230501222d5ae1a3eade09d45bcf87e4c31f88f3a6ee7212c762a79a9f426c4c637aa180c365bb246cf163c45f927ff2103ef58a3ff4b8546da79dff51f42de69705a28bac9f7705abb46e47bff0c0000ffff17e2d65d48100000"
	tmp.Length = 4168
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/purppura.js"
	tmp.Contents = "1f8b08000000000004ffec57dd6ee33613bdf7531c708395f4d9b1bcf87ad3f50f903628d002dd024def6ca36044da62a3882a358ad7d8f8dd8b21fd23c74990748ba217992b9b2372e69c99c391164d9991b1256ab9d02e26fd9992ce970ee03435ae042f74e0adef7455c84cc7e9fb74d983782f6faba1484ebca3e02de831e72438978f3a4570fed9d847dd5170bf1bfcffdba148869d4da7b34fbfa99424fdbb2cb4a33a0e1080b3fe52d34f57bf7c8a21527da74baa057ad8ed8aa124492460c0c1eea4438e31be6c865b60403e8d9c34b556d11c630cdaeb952e952997a70e99dd94765568b5dc6fdbef4b537c5f68e9601b0792d785865da0923521e488735cafe1f4adbd33e512b228e0ecaa865db4cfa05c87dd3de8cf99ae080bebfceac2b89a10af7293e530b55fcbb554da9d3bbb4afafb537843cc906ff41aa644dea682ed2c16ef04badedf85d8f22b92fec2942a16e43e2e291e2422e9fb6c759c1c68db1c109ff5b5ccf240779bfe1bbdeee14e16be04fbacd8d214df35b7954f3db34d494c11e50c665d79be42a97b1e83a7c250a10fd0d818594d929a1a630e338daefcbf687ec8922d9f86a7b8bcfbdfdd0f2719fdc6f14d8ddc3a6714ceb1d258c9924016756e57d08672ed2056a628e07b069244ef21325170ad4b4b6661b4f28fc03a5ec8b9de4a6fdb0ab60cd0d6d503646681780f6d0ce1632971ca231bd3a030c6cf92f2beb34da9e2c0c6a76d061714cdf13f7c180c063b391f8c7756d2d5fa87c24a8a55abc23be3474abdc2a5241d4301c9f0983bb6348551b8476d1b9769fed15cffa133c23d8ee9b887f457d2116b6c8c83308618914356c8ba1ecf445698ec662626235213eed43b594ca31f2fa339f7eb28e5d5e362b311bafe98d6962b9fd6cbb7a1bbbd327dc0ab80259a272f0fabf88c17a43892c89d5e8c67226ddd2a29830dc119ad8f3b131399dd8c52391127e405c8681d96f11df4c431def7dc4181d947ab7ca194ef5a2fd0933c76374aa06d2b4726edc1dd12d1b555eb28e9cb8ad51053f26c30a5499aa23ed63fdb6be39d24ccb64bc2775e4deb428f674299ba2ae4fa63694b3ddc76a0a765440a992dea4a96e399f8865dd5a1372f7da6bb46abb65b527293fd2c0db6394ae5a1e25badf094ee9f515c50587d1481ed3f25b1af5258abd1fde5f844a37b1f37faab95e14bfda6807f5a01472df9bc04f85e30e5f2b553ef577e977b1b79d3578fbcafd2238fabbf31f1fc207a42bbdec7da7d93e2bf308c8e9599a6b8d2e4e7fc232fdeca66cdad2ea9ef7dfcca76e13fcc30e536687f4f75217c75f3e9c34fa696672bf43033e7a2f572b949861d00d824c3cea6f357000000ffffbe40f640cd0e0000"
	tmp.Length = 3789
	RESOURCES = append(RESOURCES, tmp)

}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	for _, entry := range RESOURCES {
		//
		// We found the file contents.
		//
		if entry.Filename == path {
			var raw bytes.Buffer
			var err error

			// Decode the data.
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(in))
			if err != nil {
				return nil, err
			}
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				return nil, err
			}
			_, err = raw.Write(data)
			if err != nil {
				return nil, err
			}

			// Return it.
			return raw.Bytes(), nil
		}
	}
	return nil, errors.New("Failed to find resource")
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
