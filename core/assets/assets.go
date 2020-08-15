package assets

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"io"
	"sort"
)

// AssetNames returns a list of all assets
func AssetNames() []string {
	an := make([]string, len(data))
	i := 0
	for k := range data {
		an[i] = k
		i++
	}
	sort.Strings(an)
	return an
}

// Get returns an asset by name
func Get(an string) ([]byte, bool) {
	if d, ok := data[an]; ok {
		return d, true
	}
	return nil, false
}

// MustGet returns an asset by name or explodes
func MustGet(an string) []byte {
	if r, ok := Get(an); ok {
		return r
	}
	panic(errors.New("could not find asset: " + an))
}

func decompress(s string) []byte {
	b, _ := base64.StdEncoding.DecodeString(s)
	r, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}
	defer r.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.Bytes()
}

var data = make(map[string][]byte, 14)

func init() {
	data["core/00_builtins.scm"] = decompress("H4sIAAAAAAAA/3SUwZLTMAyG732K7I0u42G4EpZ3cRwlK3BkV3JCw9MzLW2HStrj//2ObP3SpO/7vosZulQYvnXDirkhHQ6fRpjCTXXh+Kxf3hR4VfqL0p+V/q61Lqj1D63fjuqNsVag8Whg3hVLkTUZDXmPpFEh0WhlBmqh4QLKGmECXRROCkzI0hSbgWRfDDTHio48xz97EDgZDjS3dw3RXLzExEWzogMlU6oMPyHpYpXLgqJD4egwiPoOhlQ2kx6DeTLDBmwqSuMXizQxGW+QWuHj4dD33RAFurZXkOczKMHbKJQwlJLB7AxKcJJGCbQug+kQJdSIDpXGSLPD92Uo2fJHK87jgUbv4tiKDgQlODt/pSs10FNDCbDU5mSDNMLZO/8L9t+FHSOXFJ2+vBW94lq98uTNg2AOSJM13D2+jKSI/4G/4yjhtMaME3pPYpCSN9+57DIO2Sko5s9xgRUSXlL651yz6WSnFs/htJYGd+d2sBtgxnsgdzZCypHh1eAJyVDYHkO5I5wUyHEZxqjh4/f1HwkMSdFrD3Cukcbw9WNPObdmn1iNrQHT8fA3AAD//zTwK9PfBgAA")
	data["core/01_basics.scm"] = decompress("H4sIAAAAAAAA/7RW3XKzNhC95yk2yUVECv6cXsbTGb+Hx/NFBuFoKiQiyZn67TurFSCw7F6Vm8Ts2b9zVit2u90OuBLQGCs+4MSdbFxRsFZ0Uos3eBuMq6Xu3oD9gvfNFrabbVkmdi3Ok72eATPC24uAF5/4dFw5AS9d8kpLBa8M/XY7cFft+T/198V4qc9gxfdFWuFA+tmhMbrh/qkAYIr3p5ZDY5RyBQC+Eb62ooEDoWqptbDBND+jGwt+8CV4W64gAEx2wKSrRT/4K2W4BQUgHwZ1BSWdvxNqLAzYoQPWSet8DHjMYuk5WGBWTFA45gOvK+3uwQI0ZQXsg3pzePzlYlFdCfaR5yoOHwah25BvZKAr6Rk5WOYiaQ5HRCQTpfhJKFS+5401UfKopua9gM7Yfqzrc56GKhgrtB4h/FgFlq7G7mpja1QyGS72w9VFUMyR6kBE8h4fHPZi1sSFOBFEuXAEDj3/W9Rz9UkeHHqRkMpSVIrELtyC+3m6Gm7Jnp0tnKmmnRBrAdP2RibyAxWLjfnKcTRCbJyMRz4d2kl5pHlUgP7W1GMRCbjHVZDzZNrrnOhzihT0LRb04NjQ+ylGtQ/+JYlzW0gyYf9PHSFgtpxkMiktcOeE9TW3ZxT+wEpcmcfwb6P4xY1zyCyXTsBzggcCOOgvzsNJwMClFe0zuUdjEUuXHVRsdNimKrI0ZLWnLRAXFOHLheqxkjnaOyp+01hsP5zSQMsmDmcxj6MyDVfBmpzsrKLkmqLIQKd/NN8W8SS+6aTCBhrTDyH9J9PGAxPfUJGt2pMtI49uoyy4BRa6RHZyZId9wHX7suD8WKTHEc3Lo851O/E/Mb+AoE+W6yf03Uxec48UM9Fx5Wnsnan7r+6MfdScscve1r+Zsbet5jszNttYCHC/L9aKRnGLsuOI0NRNl0Y0rjZtcvJvvgiWV3EWgs8ry+5H2qFh6b5yJX7F/G/jjUnx7t+6UzfEWIRPV+wnO4mz1MjoCJwgGWLwF2ym5UZ60kEIWz8YKDJbrqdRuvnsrq5KuL1l4j05ejtv4VnqH65kGxfyBzxTzuVVnE2Q7OB9fjvXN2t3ooAU/BF6jsr+AtabNu6IP0vYrvGmbR/A39fw8GU842UXseF9+Ehee9CXc8aFDOi4dNFNgv4jgpeFtCLF1FmMNrUZYAO/l9Fbob301+g+RpkBKK7n2qtrkiFy/ntSez11H+C5PQsPvfBfBncVjhSdZ4bvq2ivCFBCtQ+Isvg3AAD//w1vgQjODAAA")
	data["core/02_binding.scm"] = decompress("H4sIAAAAAAAA/4RS0W6DMAx85yusPpkJpm6PRZP2HwhpAdwuWghTklbr309N4kBQ1/mlFT6ffXdpmqYBoQiG2dABeqlHqU9FgSMdpSZAaev4sR6UOFuC8FMWACj06BEXGtxsVh1f+AavgIr0yX1yb2lKW6t5EAow8u7LsixXmyfxRbzbwkWoc1iqyNWGBmiX0yyzBmolpn4UgGkm687mnqq7WI9nkUpa9zcs46bp213/wWbcm2PwKE1a9pAheZmMQkPLqK+uCLusJeNqYZJd2VyYALTOwE7qi1By5PdwgF12DMrjOveNULznFB+UAq4nMZjZx/kE6Ypn6OfxyjlD22/fAf9hVTcUtvErG9fHbqhW04/zYG9NX0LHd8VxLSYCQGbZZ+MArT9+6b+U3Upt9CJkfluVB/YRhFR+R+WZOqjeWeZj4Db64FblBUWOUMVvAAAA//9KoHJsyQMAAA==")
	data["core/03_branching.scm"] = decompress("H4sIAAAAAAAA/5xWwW7jOgy8+yv4cqlc2MXbPTZYoP8RFKhi04kAWU4lJZv8/YKUrFqOk3RXhx7EmWhIc8iu1+s1SI3QDBZfYWulafbK7IpCtNgpg2AGD8rVndQOy3Rd97KxAxyNRucKgI3w6HwJX+dJlO8pAH6PJkY/hOqg4tsnUULFoRkUkF6bQiu6GcFXOn7v0SyoyDV0g+3LmYKKLgk3Bb4w1GXviy3ulIHqLYaYsqijNoP/Jy1cDQ480ELAmZ4FLc1gWv6dMr3daHl0SN8LhEY/XoDolHUeYpixBJHOofW1tDtXxCyENC0I5eoTNn6wkVIWkw8P4hf8pN83O78fAWWCCOctrJQ5Sa1aFhkxr7BK6CiANIoN5wwgotj/R33hbELPpvAPKgXMWy0TWAVKLpqVVG/C4qQSdKiy6zU00iFY/Dwqiw6UTwYRvTz8B93RNNAMWtPjIviCC4X9wV9SZPqcA8GssfoEKac/F7Xw/ZRMvXf1uQULxPPBwguLdaxEo68tNiA2J6nHx3do3KWH1Unq1dSz6WQVhs3BYltrRTbQst+2EoSmjhwsVYwFEwRI5wI1vjpSz0TFT6hIUXUuy69Xcy5l8ZxdjT9SzPSGJrdSOYSVGUL+0AxH3cIWoZe+2WO7momD3BKz8w2HzBkLhrmCPPbPnHLXTnNw7i5qiVvuWkp3wWxw1Q75ufbfQ4bBM7+RWe39VkrBxPSH6sZdyJP1Xs1CRl9t+xeEgL0LXpofc4iQh4O+8Bd4Bko4jJKY5UdoruAA8myqV04MPl5wu+pqjWHNbJVpldmFFVukklEcqjHIi3Oy56aksGy/P/Mj987QHxG3pn5UF2HUmomRTX4aUpC0zgd/3Msxms/9UF7i8/jPiNxPIUTZ838W5c1tPi8zL+ibZZ6s7/HuBbZDe7lJSYucUaTiTwAAAP//E3fZ7ZQJAAA=")
	data["core/04_predicates.scm"] = decompress("H4sIAAAAAAAA/7SVwc6bMAzH7zyF5x4WpDLtvGrag1STvgAGZQsJTUI39vSTCaWsrXfaONiG/Bz8N256Op1OoC1B4wN9gjFQaxqdKBaFspSqQA2oMz+u9DjauYD7pawe6laD6ibXgA59LPfLAMp0oEysaBjT/ArgK4WJnp+qyVmKcd1bdSbElHcoodM2vkjZUu/VQs4OtCWXZfm12OWeW+qMo2r0UZbm9ECP0iwlOPNqxaug4jyAiiksMOAX5Bc9VPam8tvguEvc3nSADxAOLzoE6riTdMyaGu8idAfOkFU56v+RKnyH/13ZlnDVdqISlPMJVNabH5Wv9vhzw6e+lNyYtbxq0E3w29027fDQkDdVU28cHNV9PB6Z27Uxjvodw7UWt67wj8BN1q4yuCC65Bt4r57QXxT8Dv28kh933K52E7l/5FoKgDnCUuTsvEB2FpnkB0C2ElF7b0k7wDWQuOVTIFuZmFyilqElkLh8gODiROZKDpCtRCznBuDiJMa4ln5yRWsgcd9p/uFDC7gGEmdNTIBsRcI32gIuTmLy2OLiZGYcufDsJcrxd3PyN3PUV8Z1gGxFiEcZ2crEUPM4Zi9Rvv5GTQLMXqTaFtC3oqRRmwDIViR8zKL+omkMfjA8HmsgcZdJW9MZ7vQWSmyg6O2V0Vskk1cK0dSWmL3FEh3pAhjpIq6P1BgeqjUQuRSM6wGzF6l5qD1vtniJ4r9xQLYScaUm+QCYvUQtZx+yxbL4HQAA//9NhdK4nggAAA==")
	data["core/05_concurrency.scm"] = decompress("H4sIAAAAAAAA/3xTwY6bMBC98xVPysVelUq99JAc2v9ASDVmSJGMJzWQbvbrV9iYtQkKl9gz8/zevJlcLpcLlCFodnSGZqtn58jqR1GIlrreUjko7RjiyviOhtuHLIA/4spvEEYNTasgJL799ikpn3DdPM2OMqyhCdXN8dCPdIJYT0fP1QWWbyGPZScpQzC57zlbMuqRUb7g+IJDdOw04a7MTAtO9N3G++sr7CVlt3B5FmLUR67DK/Og4MKSKR2Ns5lOEA1dexuF1ZFp1Vw60nA0srkTRMBEfuy1PqV9yR69y++jqTh5NNsrWXJqepruG0Sl/yp7AsTyKyNDnRBW2rCf/9kf4AEyq6ChnxbdZ38IBajlthSxNnj52salLDAmXW/NrU+eR/q38qTtrksz3tR/WwCV6GarI8ZHsUTw46enC3kMDb+XY/9BB5VbDpZLvh3CMLDtJ3YRnhhryQRnE5I6G2Y1qN40/J70ZMnIfARANZJtyWUe+7p6cynx2a+Q5ju59I+0yg58+50KX+zkMBkkyFoWnwEAAP//bjriE48EAAA=")
	data["core/06_sequences.scm"] = decompress("H4sIAAAAAAAA/6RWTVPjMAy951eovWDvThnOdAZ+SKcH11HATHCK7XSXf78T+SO2k1Jm6QViPUnvPVlp9/v9HkSPIAeDj2Cd0K0wLVj8GFFLtE3DWuyURmAWP+Ai+hF5A8BUB0zolo6f5/P0YRt8P7vPGIoxegr/3zHOedlgU3VYFqfjsrb/Q+UiyAhlp4rOwNZHpdB3DoR0ICyIJPARtpFhzWX3NJzeUDq4Bzn0vSVS4nzuPyEEwpMctBRuE1DLMr2ya0Xo+JslLijdYFaKhMDtMuJ8Rt3+ohDce82+EOvF+6kVO4MSPGqntEYDjLAzMtmfboLNR87KZP/k+7FOGetiDgdmMHvMakxof5D3znXIQdtbKibMf2ooU7WtqHuCP+T/tkKf1okce/ZFmtnU/hPi9DynsnYB9PaswAoSPeoX97pJnQr3fLDw72zwUrgnh1E7rLhS9DeBY4vMEKjMrzMpWPYml4PjSktPo87xh7PnD7lS7V43DdV1pO1gpdBN2ZKUR6GDhXdlrdIvC25fEU+AJyrxsBYnzNS/FNainHL4tb4p1V/EhaEpHvM5PzaZMtIdfKDPIUldzEXpFv+ujZUQOsyzTp2FJQ+Tqzy+i7dUG4bRwdDBaRh1a7c8kS14QYudGHv3A35rJb7mGRM4P1bLIqLt4SbBwZb3O/j95xX1dAOq+c+8aNblZtSLKay7tpbCuq+Xkr6Or25mwBUOFjSvKytmsND3HZUL+GKVb70XHLBDVyzBcXVRDuV747jCNHPSQHfzhVL+TNmFG2OwHeX0S+bAulFLUFq5nHsxOw+O06uht+SzMn1uV7wSctlVgQmczubux2amnzO/QiXQAMLnra90psLTNmX32+AFjcXsik/N/Kk69Vh9+QV4fUg06Iv0jvF0xf4FAAD//5liNqbMCgAA")
	data["core/07_lazy-seq.scm"] = decompress("H4sIAAAAAAAA/6RXy27bOhPe+ykm2fxkf6uwt02R9D0EA6GlkS2UohSSbuM+/QEv4kWifAocbRJr7t98M6ReXl5egHGEZpT4DTj7cweFHzcUDardjrTY9QKrgTVyBGLElcIP+Arnsb3THcB7ePnFyIdzy4BQ2P+wCpQGH0A0+2ni3ISGZuTcWBNvUklswMirXgiUQHI184Q4/jcA6TsgTLRAXr3bAw3C8JAnHCZ9f3POcgXSjEIB6XqpfDCX5ZxFi43zTIFInHUWTv5HqH+VJL2su/p97TnCJLGNxWcVkd9XFBVHDbUCYkC2eqdQvpF0IVsVJN4UiHXeqSUGrkhjuUzE1aSof5KUWzlOy1bl2aaNM9rbjfON2ugRSY038M4M0lcp4DH9ytNwYlL3uh/FDqBeJ0aCPPgJGqddYqM0Tpnhkoke/9CzXQF/I61eea/0ehAKrC2kZ9NwnYkZmee0Ll0ycUFbwpyzfQMHw1Y4zvVxpvRSwbybO2b/P1A4QnWks5HjX2qaKFthUo93G23guJZGy3IcW28A38xBM0y7BcQt1ETcOH9zqTnYNROa30HLGybzEp6afHdYHigAvC7FJYtvyBXOv76vLGIUi0ozTLDCK6WFE8bF6hEh//cSW3kCQs6V+lTs/sAmh+BNNBlx06Ed2FSpXlw4mlQy1q5X7b+SHJJtaqKmO7VIb2uQ5pDu11hmOo2hHPhq/6itsiYmGefIXWHqLytj0+Sp8mYzg/TgUKUqLBdJ3fnh/oWNHqWzjeUrWmKe75/XDrUrutA4lXeDQ9plbGHpqEMzlC4zGJ2BxW7OKV33Xc81Ssj48mDfO/W48bNd/5gmEbKEIGV8DDiRFCUk7Ij58osu/IzlBcoNQpKsLJnAtzpg/HWoG2WFrLmaG1OFn5NUi4uR0k9WC/ZRo3QzakbRMJ3SOsPbiVO81YNbUd6A9XWg0AH1Ny1Q2z0wwR62wMfa0ImaWZnGzkbvqG3Hpv9olbB+ntqI8p9+yiB242Pmzx7La4OBTaYtljVrQ9+0xE2YH7XBliJR3Il2NuF+YnXuRduLi4qqvjeEKYVSV0xe1Nz8eUOdKTwzDRzNOTEKBBPLe4JegcSPWy+xfXYYmohfgNT62ofWnHMK1AI/tYf/TPOTsFb3AYBY68NSZlL2smNCGEsTn61xnTbz3e7A8PGwV/chfEBYbj3B3rilKyOLf27nxs1mP09a7oOWetMwqVH1TFSTHNtbk8+ih0vdh+qCAmLITwrkgsLgQZSW8NxMz/BJadziCTg1R12p++Av4satuOjrPFsWhDmEuwbQ7Oiw9r8YV141PXIMtYN7z8BifEOK9C46+1yeO6X8zbr7L/kb+wf5B/dzIaaEpf0y/9nnZv6OM+92yPYRAcuS6ND62v+YU3Ak+ScAAP//vgeltRwPAAA=")
	data["core/08_threading.scm"] = decompress("H4sIAAAAAAAA/+xVy27bMBC8+yumzqHLQiySHmvU9n8IBsLYa0coRTkk5bZ/X/ClyI7ipodegvIigNzHzOyQWiwWCyjN2HaWv8I/Wla7xhxmM9rxvjEMSnvS8ZNc6sZ5eGUP7MUMoN5odg4U9lejg7LoIqGsvCOGNrJVW9tBLmdATSelexaIn83zDj5j39nW5Uqk2X8C1XoCI+0b63wOF5tRa6Del2MtML0uEizI8rV4bAZ29ySXoGqPKmGu1lagWqcCGY/YTDB/H9Qz97XN/N/C3XUtv4vJa/aoT0rfZPLjYvTjkY00nQeZXutV4HgjZlMlKQkSlUzVpjz0upL/pbyU8tmVMfRNWgYIcomHxoQHcSwbKefYeqnswcWu1DiZ4+RWq95xSROYlwL5oO2dxwND4agayzuceOs7O0+FPnB79L9WudcZlbny0KycR2c4v9S8i5FoHCw/9aHePCIM8lHtPB/xF4OtDf/0wFib64OtjWoZoMLxdnq2o4TkwSHh7pUEkXXdg7IiAdrzjLNBqtg/OwRV4Hs9ZmySPOASEJI3qNaxUVijf2CrvrPcdmZXxpvlt0ls1T7sFCidZeiXHgEoTXqVnZCZz1/4YmwIgL7hS4BqDv6xZIoXWclL89w7Tf8Y3JVB4fbs/tbRNeXsTgw3754GNielRZxAlQpVhXO4Q6hChfi6iHOx8u2JYv3pIUr9y1MU/5xpTiOo1ZpadZwYwUe5FEOFifubEPxrCOcYfgcAAP//ImobVk8JAAA=")
	data["core/09_functions.scm"] = decompress("H4sIAAAAAAAA/4RUTW/bMAy951cQuZQqku6+AGt/wLDDdjQMVLXpQINNJZIMzP9+kGTJsuelOiXG48d7fOTlcrmA7AkabegrdCM3Tmm2hwO21Cmm8yAbowF7ch3Dh+JW8dXCC3zodhIHAMReDh+tPBtq4CaNpXNGoR4dKPYw/1B1gJbuxScImZ8Bq46/K+sAO2VsCILl1QdYv6rjX9MACR1jxQP0DzlQQqOhJUbsoX9Shq3BUIsCjtJaMu4szdWusqDkFrBX1r2m0G0VANQGkO4QqTwtKu5gA37Byp6+FPi95L1uZO+re+YbxDEPaBitg0azk4qB5UDtYoHjiup2svJ2I27BDzjJG5nWYtZMsSjqvvs5B4+cvAXO37w6PlzA6S14KYKrOptMiMWFka33l1OyPwBU6BsVod06/4cXCDMOI5mL703Jt99PrxBzHOe0YOg+KkMWZJYhqTC7HDy1nVQBn35HbZ4BG822aGgJFkKIWmyXrNHDLVARoFpip9xUP6Q6cpNZzmsU6qTWrsR2GuDoPx7zQq0dXylmMtk2fipPBaXQ7SYiMtKjI5M3MLSyYDYRgXgMSBu1CSg2a741VXlYUrHUUpSyFGBt/3RpHsMCdMXjc/xKqUWiT9CBc5H8v91s0q7u03sy4emfi3OaaZSjiYNdpN83Hf4e/7jSTF5+iD7aGCi2st9FdXrDQd7ynmAnNm6aCc2Far/vfwMAAP//h6Qqyn4GAAA=")
	data["core/10_exceptions.scm"] = decompress("H4sIAAAAAAAA/6xWzXLjNgy++ymwziFUJ2qbHpNpk/fgeCa0BDnq0pQCytl1O333Dv8kiyJl77Q6JLYFgB8+AB/x/Pz8DEIiVB3hE+D3Cvuh7ZTebJjEoVHAmRTHfS1KwgpaXVZCSmD6fIRKipPGYgPTw4SqgcmuEvIF9Pk4e+ksZKuHl5Sve40fxg9Y05IegllRFJvNFCHCM1Tv5b5VdasOwJqOjglIn1gNHb3A4rWz+R1+AyZRHYZ3Z5IC7rKyJ8DjdUzAHHroBWmsE6ACm/fOPsvJMks1BHODZOnxRePHC7CnveyqrwHAGuKmVULK8w9gDh5Z1AGDSy5g+B+w4veeVoF+6egHyU0klAMw0LkUfY+GDHc4fMXzt47q5EBUnfrTgwQeDN2vLHz1OQT/3drRPeF/P1uH5okhXDnbGAXm9fysk5KoNTBbzJRFeP72hb5nBfjesB8D/Xz3T8rNqNFPwHgzlwZd7FLW/uEEjPDCGFate2BTmlQA7JIZOGbrXCQ+zis00BfG/LJwfci6KWZuWWQ2XmDHRmQXPdhPzM0DrsezM+ThLfC5At2M7wmlxhk9JFqNsD0KacQSa9s9TsI82K3ptLVec9Y9Yd1WYkBg5iMgURlfK6EzJLDhnVDUpcaP8g9zz4BxSnbI1EdynmYyz7GPYuNFh7jpavx/Dxfollz3JNR0a+h0rkJrpKEUdNDRwbPBg63rMR9TdQP0oiWstxF14q+zoSuK5RkNIjEfuCxLAPxTkA3AxssJfi3yvAJw24jg6b0cVqNFsYtj1a4QcC8k/iJxWBloh8fzuDb5wBshrbSZ+DbyHg+tAoMugcOjmWoXahapTSjhLeUfQ6zWP1Owb++obAtkpXe9pjlyuJ26eUEfi6QucpYa2+zULjn0vRqnf0Un9l19Tl9Jb94UWDHW11X14XXM+2pdMtedxAE4EgE7oDIjvkWibczjBOEBiRLLR+YeeXhlou/lGWynJ/rsgqTESjVqMh/ohGDO3t0kt+MWmNqqTMqMu7sB4n0tOd+csOo+keIlMG1cSRTq1AMbLzRvntLYBWs8CKANklrfLXzb9ndBbsIQOJ9Y2/N3Hjfud/GGNeIOKuDC7tIC8sZqbJAy0jK17gNb1scenyx85Pvq8i2KYrfJMOZrtAz25ikjrCxhrpQOjp062wAzgGOw/FZmhsaFuxuvhutPPhyhDuEebw+XpI61DXh0dn8xoQv3N0+fJ+FtlBb3Q8J+tibdMyts1sw0QquwPIqKOjvt8PNcxxMtcLGpBlMz4P8GAAD//3C8y43QDwAA")
	data["core/11_io.scm"] = decompress("H4sIAAAAAAAA/9xTzW7bPBC86ynm03chBStNe7TROu+hCChNrW2iFKmQdAy/fUHq33EL5FodbGjJmZ2dHe12ux2EJkjraAv1xWYZa0hq4QiFMgUKewkFCnKu4OnsqAyh8J2QVADIX33O57qhq1aGCuSvJl8AWOfKVnTlVYVzaS5a43gxEp7eeAawVnRgWrSHRoC9C30hDqaOYPHqHkMl/YElZF/iEZ2eyLTuhyccrWt9aqApoGpF11HzQIsP7r/hcp0I2fVMBszTG3rQ1CcebgdXtlenQhSknA/jxUHS3xmO1pUk5BkVaWrBHM0E9eLin5uOG4h4Hp/V7GY1vOg6fUPnFqV7tnFvdzzKhM/Z+E+5qEzQj52MvnzWTKEcNeU7yWDdGOnEaRqwD+VebgTt56zPbMl025EppbaeFnRpRRJsuzwY9hG/qDTCHpJDwtjSdkvishXS2QU9Dso0ypw8nnCwza1X7D25UAp38tkkczHbiOHIHxC1Fx9wIAj8olvZf9M9Mk/s0pom0VbsexzHnMJ5Qfk8uPOTHeikDDYvSVedDZgfj0DfJlC0h1UbNsl55li8feUfQ7OOUpV8/X9K1P0mVtS8ngPLgrst0ztbU21e+uz2v5NsXo/TrVKvjND6BtYLiamtOc9+BwAA//8zmRAOywUAAA==")
	data["core/12_os.scm"] = decompress("H4sIAAAAAAAA/3zPwUrEMBAG4Hue4qe9zASr9Wo9+B7LgrGdLoF0KpPU55c2i0UQcwqZ7//DDMMwICTBuJq8YM3O0SRjCibwol8ePtgtez7e56jSLWG0FVTiInjEvNqS2QHvlKR40CWXYKUFaNzMREu3S8bV4edcTPKWSgv6kFtUPLzVml9GdGr3y38902bVdDh4/ZuvXBHFGfSKQz33fc9nlj4taklah43m5q8hPZ3hPY9myQ3f5X0JZvcdAAD//+SGnjxGAQAA")
	data["core/13_alist.scm"] = decompress("H4sIAAAAAAAA/2xQS26FMAzcc4pZOot3AqReBLEwwahpTQJJqorbV4QUKGpWkWc8H7ctWAU2RMELrC7lpqFRJucFxCkFi0/ZYIOqaQD6fhcPSrKeI4BUMjpRmUGTiykfWF9AgL68Skog9uOxulNNRf88krXYVZnCM0/mLRZF+XV7suquudfZ673ewvAhNl+VeFl0wzG9RGjmxXIGKc/DyKAiiO4erfqXf/9PIZqcZok4z7Xn+QkAAP//QzCIXnMBAAA=")
}