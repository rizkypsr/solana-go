package serum

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "markets.json",
		FileModTime: time.Unix(1604944659, 0),

		Content: string("[\n    {\n        \"address\": \"EmCzMQfXMgNHcnRoFwAdPe1i2SuiSzMj1mx6wu3KN2uA\",\n        \"name\": \"ALEPH/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"B37pZmwrwXHjpgvd9hHDAx1yeDsNevTnbbrN9W12BoGK\",\n        \"name\": \"ALEPH/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"8AcVjMG2LTbpkjNoyq8RwysokqZunkjy3d5JDzxC6BJa\",\n        \"name\": \"BTC/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"CAgAeMD7quTdnr6RPa7JySQpjf3irAmefYNdTb6anemq\",\n        \"name\": \"BTC/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"HfCZdJ1wfsWKfYP2qyWdXTT5PWAGWFctzFjLH48U1Hsd\",\n        \"name\": \"ETH/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"ASKiV944nKg1W9vsf7hf3fTsjawK6DwLwrnB2LH9n61c\",\n        \"name\": \"ETH/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"8mDuvJJSgoodovMRYArtVVYBbixWYdGzR47GPrRT65YJ\",\n        \"name\": \"SOL/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"BJ3jrUzddfuSrZHXSCxMUUQsjKEyLmuuyZebkcaFp2fg\"\n    },\n    {\n        \"address\": \"Cdp72gDcYMCLLk3aDkPxjeiirKoFqK38ECm8Ywvk94Wi\",\n        \"name\": \"SOL/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"BJ3jrUzddfuSrZHXSCxMUUQsjKEyLmuuyZebkcaFp2fg\"\n    },\n    {\n        \"address\": \"HARFLhSq8nECZk4DVFKvzqXMNMA9a3hjvridGMFizeLa\",\n        \"name\": \"SRM/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"68J6nkWToik6oM9rTatKSR5ibVSykAtzftBUEAvpRsys\",\n        \"name\": \"SRM/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"DzFjazak6EKHnaB2w6qSsArnj28CV1TKd2Smcj9fqtHW\",\n        \"name\": \"SUSHI/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"9wDmxsfwaDb2ysmZpBLzxKzoWrF1zHzBN7PV5EmJe19R\",\n        \"name\": \"SUSHI/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"GuvWMATdEV6DExWnXncPYEzn4ePWYkvGdC8pu8gsn7m7\",\n        \"name\": \"SXP/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"GbQSffne1NcJbS4jsewZEpRGYVR4RNnuVUN8Ht6vAGb6\",\n        \"name\": \"SXP/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"H4snTKK9adiU15gP22ErfZYtro3aqR9BTMXiH3AwiUTQ\",\n        \"name\": \"MSRM/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"7kgkDyW7dmyMeP8KFXzbcUZz1R2WHsovDZ7n3ihZuNDS\",\n        \"name\": \"MSRM/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"DHDdghmkBhEpReno3tbzBPtsxCt6P3KrMzZvxavTktJt\",\n        \"name\": \"FTT/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"FZqrBXz7ADGsmDf1TM9YgysPUfvtG8rJiNUrqDpHc9Au\",\n        \"name\": \"FTT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"5zu5bTZZvqESAAgFsr12CUMxdQvMrvU9CgvC1GW8vJdf\",\n        \"name\": \"YFI/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"FJg9FUtbN3fg3YFbMCFiZKjGh5Bn4gtzxZmtxFzmz9kT\",\n        \"name\": \"YFI/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"F5xschQBMpu1gD2q1babYEAVJHR1buj1YazLiXyQNqSW\",\n        \"name\": \"LINK/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"7GZ59DMgJ7D6dfoJTpszPayTRyua9jwcaGJXaRMMF1my\",\n        \"name\": \"LINK/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"BAbc9baz4hV1hnYjWSJ6cZDRjfvziWbYGQu9UFkcdUmx\",\n        \"name\": \"HGET/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"uPNcBgFhrLW3FtvyYYbBUi53BBEQf9e4NPgwxaLu5Hn\",\n        \"name\": \"HGET/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"3puWJFZyCso14EdxhywjD7xqyTarpsULx483mzvqxQRW\",\n        \"name\": \"CREAM/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"EBxJWA2nLV57ZntbjizxH527ZjPNLT5cpUHMnY5k3oq\",\n        \"name\": \"CREAM/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"8Ae7Uhigx8k4fKdJG7irdPCVDZLvWsJfeTH2t5fr3TVD\",\n        \"name\": \"UBXT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"46VdEkj4MJwZinwVb3Y7DUDpVXLNb9YW7P2waKU3vCqr\",\n        \"name\": \"UBXT/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"Hze5AUX4Qp1cTujiJ4CsAMRGn4g6ZpgXsmptFn3xxhWg\",\n        \"name\": \"HNT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"Hc22rHKrhbrZBaQMmhJvPTkp1yDr31PDusU8wKoqFSZV\",\n        \"name\": \"HNT/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"FJq4HX3bUSgF3yQZ8ADALtJYfAyr9fz36SNG18hc3dgF\",\n        \"name\": \"FRONT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"HFoca5HKwiTPpw9iUY5iXWqzkXdu88dS7YrpSvt2uhyF\",\n        \"name\": \"FRONT/USDT\",\n        \"deprecated\": true,\n        \"programId\": \"4ckmDgGdxQoPDLUkDT3vHgSAkzA3QRdNq5ywwY4sUSJn\"\n    },\n    {\n        \"address\": \"5xnYnWca2bFwC6cPufpdsCbDJhMjYCC59YgwoZHEfiee\",\n        \"name\": \"ALEPH/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"BZMuoQ2i2noNUXMdrRDivc7MwjGspNJTCfZkdHMwK18T\",\n        \"name\": \"ALEPH/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"EXnGBBSamqzd3uxEdRLUiYzjJkTwQyorAaFXdfteuGXe\",\n        \"name\": \"BTC/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"5LgJphS6D5zXwUVPU7eCryDBkyta3AidrJ5vjNU6BcGW\",\n        \"name\": \"BTC/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"5abZGhrELnUnfM9ZUnvK6XJPoBU5eShZwfFPkdhAC7o\",\n        \"name\": \"ETH/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"DmEDKZPXXkWgaYiKgWws2ZXWWKCh41eryDPRVD4zKnD9\",\n        \"name\": \"ETH/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932\",\n        \"name\": \"SOL/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"EBFTQNg2QjyxV7WDDenoLbfLLXLcbSz6w1YrdTCGPWT5\",\n        \"name\": \"SOL/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"H3APNWA8bZW2gLMSq5sRL41JSMmEJ648AqoEdDgLcdvB\",\n        \"name\": \"SRM/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"8YmQZRXGizZXYPCDmxgjwB8X8XN4PZG7MMwNg76iAmPZ\",\n        \"name\": \"SRM/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"4uZTPc72sCDcVRfKKii67dTPm2Xe4ri3TYnGcUQrtnU9\",\n        \"name\": \"SUSHI/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"9vFuX2BizwinWjkZLQTmThDcNMFEcY3wVXYuqnRQtcD\",\n        \"name\": \"SUSHI/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"33GHmwG9woY95JuWNi74Aa8uKvysSXxif9P1EwwkrCRz\",\n        \"name\": \"SXP/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"C5NReXAeQhfjiDCGPFj1UUmDxDqF8v2CUVKoYuQqb4eW\",\n        \"name\": \"SXP/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"FUaF58sDrgbqakHTR8RUwRLauSofRTjqyCsqThFPh6YM\",\n        \"name\": \"MSRM/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"58H7ZRmiyWtsrz2sQGz1qQCMW6n7447xhNNehUSQGPj5\",\n        \"name\": \"MSRM/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"5NqjQVXLuLSDnsnQMfWp3rF9gbWDusWG4B1Xwtk3rZ5S\",\n        \"name\": \"FTT/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"ES8skmkEeyH1BYFThd2FtyaFKhkqtwH7XWp8mXptv3vg\",\n        \"name\": \"FTT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"97NiXHUNkpYd1eb2HthSDGhaPfepuqMAV3QsZhAgb1wm\",\n        \"name\": \"YFI/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"Gw78CYLLFbgmmn4rps9KoPAnNtBQ2S1foL2Mn6Z5ZHYB\",\n        \"name\": \"YFI/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"hBswhpNyz4m5nt4KwtCA7jYXvh7VmyZ4TuuPmpaKQb1\",\n        \"name\": \"LINK/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"WjfsTPyrvUUrhGJ9hVQFubMnKDcnQS8VxSXU7L2gLcA\",\n        \"name\": \"LINK/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"GaeUpY7CT8rjoeVGjY1t3mJJDd1bdXxYWtrGSpsVFors\",\n        \"name\": \"HGET/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"2ZmB255T4FVUugpeXTFxD6Yz5GE47yTByYvqSTDUbk3G\",\n        \"name\": \"HGET/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"FGJtCDXoHLHjagP5Ht6xcUFt2rW3z8MJPe87rFKP2ZW6\",\n        \"name\": \"CREAM/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7qq9BABQvTWKZuJ5fX2PeTKX6XVtduEs9zW9WS21fSzN\",\n        \"name\": \"CREAM/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7K6MPog6LskZmyaYwqtLvRUuedoiE68nirbQ9tK3LasE\",\n        \"name\": \"UBXT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"DCHvVahuLTNWBGUtEzF5GrTdx5FRpxqEJiS6Ru1hrDfD\",\n        \"name\": \"UBXT/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"9RyozJe3bkAFfH3jmoiKHjkWCoLTxn7aBQSi6YfaV6ab\",\n        \"name\": \"HNT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"DWjJ8VHdGYBxDQYdrRBVDWkHswrgjuBFEv5pBhiRoPBz\",\n        \"name\": \"HNT/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"AGtBbGuJZiv3Ko3dfT4v6g4kCqnNc9DXfoGLe5HpjmWx\",\n        \"name\": \"FRONT/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"56eqxJYzPigm4FkigiBdsfebjMgAbKNh24E7oiKLBtye\",\n        \"name\": \"FRONT/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"AA1HSrsMcRNzjaQfRMTNarHR9B7e4U79LJ2319UtiqPF\",\n        \"name\": \"AKRO/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"FQbCNSVH3RgosCPB4CJRstkLh5hXkvuXzAjQzT11oMYo\",\n        \"name\": \"AKRO/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"Fs5xtGUmJTYo8Ao75M3R3m3mVX53KMUhzfXCmyRLnp2P\",\n        \"name\": \"HXRO/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"AUAobJdffexcoJBMeyLorpShu3ZtG9VvPEPjoeTN4u5Z\",\n        \"name\": \"HXRO/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"ChKV7mxecPqFPGYJjhzowPHDiLKFWXXVujUiE3EWxFcg\",\n        \"name\": \"UNI/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"GpdYLFbKHeSeDGqsnQ4jnP7D1294iBpQcsN1VPwhoaFS\",\n        \"name\": \"UNI/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"6N3oU7ALvn2RPwdpYVzPBgQJ8njT29inBbS2tSrwx8fh\",\n        \"name\": \"KEEP/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"sxS9EdTx1UPe4j2c6Au9f1GKZXrFj5pTgNKgjGGtGdY\",\n        \"name\": \"KEEP/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"5P6dJbyKySFXMYNWiEcNQu8xPRYsehYzCeVpae9Ueqrg\",\n        \"name\": \"MATH/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"CfnnU38ACScF6pcurxSB3FLXeZmfFYunVKExeUyosu5P\",\n        \"name\": \"MATH/WUSDC\",\n        \"deprecated\": true,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7NR5GDouQYkkfppVkNhpa4HfJ2LwqUQymE3b4CYQiYHa\",\n        \"name\": \"ALEPH/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"CVfYa8RGXnuDBeGmniCcdkBwoLqVxh92xB1JqgRQx3F\",\n        \"name\": \"BTC/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"H5uzEytiByuXt964KampmuNCurNDwkVVypkym75J2DQW\",\n        \"name\": \"ETH/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7xMDbYTCqQEcK2aM9LbetGtNFJpzKdfXzLL5juaLh4GJ\",\n        \"name\": \"SOL/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"CDdR97S8y96v3To93aKvi3nCnjUrbuVSuumw8FLvbVeg\",\n        \"name\": \"SRM/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7LVJtqSrF6RudMaz5rKGTmR3F3V5TKoDcN6bnk68biYZ\",\n        \"name\": \"SUSHI/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"13vjJ8pxDMmzen26bQ5UrouX8dkXYPW1p3VLVDjxXrKR\",\n        \"name\": \"SXP/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"AwvPwwSprfDZ86beBJDNH5vocFvuw4ZbVQ6upJDbSCXZ\",\n        \"name\": \"MSRM/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"FfDb3QZUdMW2R2aqJQgzeieys4ETb3rPrFFfPSemzq7R\",\n        \"name\": \"FTT/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"4QL5AQvXdMSCVZmnKXiuMMU83Kq3LCwVfU8CyznqZELG\",\n        \"name\": \"YFI/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7JCG9TsCx3AErSV3pvhxiW4AbkKRcJ6ZAveRmJwrgQ16\",\n        \"name\": \"LINK/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"3otQFkeQ7GNUKT3i2p3aGTQKS2SAw6NLYPE5qxh3PoqZ\",\n        \"name\": \"HGET/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"2M8EBxFbLANnCoHydypL1jupnRHG782RofnvkatuKyLL\",\n        \"name\": \"CREAM/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"3UqXdFtNBZsFrFtRGAWGvy9R8H6GJR2hAyGRdYT9BgG3\",\n        \"name\": \"UBXT/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"9jiasgdYGGh34fAbBQSwkKe1dYSapXbjy2sLsYpetqFp\",\n        \"name\": \"HNT/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7oKqJhnz9b8af8Mw47dieTiuxeaHnRYYGBiqCrRpzTRD\",\n        \"name\": \"FRONT/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"F1rxD8Ns5w4WzVcTRdaJ96LG7YKaA5a25BBmM32yFP4b\",\n        \"name\": \"AKRO/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"6ToedDwjRCvrcKX7fnHSTA9uABQe1dcLK6YgS5B9M3wo\",\n        \"name\": \"HXRO/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"FURvCsDUiuUaxZ13pZqQbbfktFGWmQVTHz7tL992LQVZ\",\n        \"name\": \"UNI/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"EcfDRMrEJ3yW4SgrRyyxTPoKqAZDNSBV8EerigT7BNSS\",\n        \"name\": \"KEEP/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"2bPsJ6bZ9KDLfJ8QgSN1Eb4mRsbAiaGyHN6cJkoVLpwd\",\n        \"name\": \"MATH/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"B1GypajMh7S8zJVp6M1xMfu6zGsMgvYrt3cSn9wG7Dd6\",\n        \"name\": \"TOMO/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"H7c8FcQPJ2E5tJmpWBPSi7xCAbk8immdtUxKFRUyE4Ro\",\n        \"name\": \"TOMO/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"rPTGvVrNFYzBeTEcYnHiaWGNnkSXsWNNjUgk771LkwJ\",\n        \"name\": \"LUA/USDC\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    },\n    {\n        \"address\": \"7PSeX1AEtBY9KvgegF5rUh452VemMh7oDzFtJgH7sxMG\",\n        \"name\": \"LUA/USDT\",\n        \"deprecated\": false,\n        \"programId\": \"EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o\"\n    }\n]\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1604944659, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "markets.json"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`data`, &embedded.EmbeddedBox{
		Name: `data`,
		Time: time.Unix(1604944659, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"markets.json": file2,
		},
	})
}
