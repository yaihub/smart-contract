/**
 * Copyright (c) 2018, 2019 National Digital ID COMPANY LIMITED
 *
 * This file is part of NDID software.
 *
 * NDID is the free software: you can redistribute it and/or modify it under
 * the terms of the Affero GNU General Public License as published by the
 * Free Software Foundation, either version 3 of the License, or any later
 * version.
 *
 * NDID is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the Affero GNU General Public License for more details.
 *
 * You should have received a copy of the Affero GNU General Public License
 * along with the NDID source code. If not, see https://www.gnu.org/licenses/agpl.txt.
 *
 * Please contact info@ndid.co.th for any further questions
 *
 */

package data

import (
	"github.com/ndidplatform/smart-contract/test/utils"
	uuid "github.com/satori/go.uuid"
)

var RP1 = utils.RandStringRunes(20)
var IdP1 = utils.RandStringRunes(20)
var IdP2 = utils.RandStringRunes(20)
var IdP4 = utils.RandStringRunes(20)
var IdP5 = utils.RandStringRunes(20)
var IdP6 = utils.RandStringRunes(20)
var AS1 = utils.RandStringRunes(20)
var AS2 = utils.RandStringRunes(20)
var Proxy1 = utils.RandStringRunes(20)
var Proxy2 = utils.RandStringRunes(20)

var UserID1 = utils.RandStringRunes(20)
var UserID2 = utils.RandStringRunes(20)

var UserNamespace1 = "cid"
var UserNamespace2 = "passport"
var UserNamespace3 = "some_id"

var RequestID1 = uuid.NewV4()
var RequestID2 = uuid.NewV4()
var RequestID3 = uuid.NewV4()
var RequestID4 = uuid.NewV4()
var RequestID5 = uuid.NewV4()
var RequestID6 = uuid.NewV4()
var RequestID7 = uuid.NewV4()
var RequestID8 = uuid.NewV4()
var RequestID9 = uuid.NewV4()

var AccessorID1 = uuid.NewV4()
var AccessorID2 = uuid.NewV4()
var AccessorID3 = uuid.NewV4()
var AccessorID4 = uuid.NewV4()
var AccessorID5 = uuid.NewV4()

var ReferenceGroupCode1 = uuid.NewV4()

var ServiceID1 = utils.RandStringRunes(20)
var ServiceID2 = utils.RandStringRunes(20)

var NdidPrivK = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA30i6deo6vqxPdoxA9pUpuBag/cVwEVWO8dds5QDfu/z957zx
XUCYRxaiRWGAbOta4K5/7cxlsqI8fCvoSyAa/B7GTSc3vivK/GWUFP+sQ/Mj6C/f
gw5pxK/+olBzfzLMDEOwFRbnYtPtbWozfvceq77fEReTUdBGRLak7twxLrRPNzIu
/Gqvn5AR8urXyF4r143CgReGkXTTmOvHpHu98kCQSINFuwBB98RLFuWdVwkrHyza
GnymQu+0OR1Z+1MDIQ9WlViD1iaJhYKA6a0G0O4Nns6ISPYSh7W7fI31gWTgHUZN
5iTkLb9t27DpW9G+DXryq+Pnl5c+z7es/7T34QIDAQABAoIBAD/nq941tKx/2ppe
V/V7CZ6zc05OZN3BNBFJi9QbJO3D4dOigx4ib7Lg6n6bAkuqLK9joh+oQW8X+eG8
G1btEGwaTr0kPVMDa6xDUleUOXSVMTCyCvGSfXkaufEwv22nVzknYk0W6hCiATEw
lR6Akdmr3mIg8jwXNRVThO8MPFNWJK2TEKM+VYyRJaHrTiVnnGBaAc+6jM19xh9V
92j0O/+wN+XvOt0m41+PZxz37nKRqX0HVqo/RZJ7OwzyGtPdNMd9AlXftS8eQhhG
GopFPuEjWDjAIziC1MBtI23BFAp7cb7hkDK4p0D2ZZRrcGizA3ah/hvv0cUGaBb1
EMzJmvECgYEA//D9aUm8T62OZzNdCC0lxPx/tS3kG8tf1hMGjU1zT5JResCZCvsk
Xd1PS/62EWg5KHgz4Vn1eOApbYtPDiKOSiZAj2/pLhvRpStukC5I7ITDnE0sgDbt
I/kzfcGR8TsVZVHu9FIvoZ8WrzBTNwC1uihOpxFpVOS23365fUXP9t0CgYEA31XS
v2pjBpP5a7TFzwz7ULSBzO+sR1Qm++W7bUM79DJfhiB1piRordhxjERL1W95FX2z
3/V4bQ7ophPkfcy68RCC5Wfts/+lHQoVNlG6suEsmBsac0g7ONcsaMrEvn9JBVxE
g6bO7CRRHZ9o7nl4k0nWEbxhyraWHFbqBISQutUCgYA8czf3QUIn848Z0ujbQIaW
Mykaqt8grXVSQ6Ydg7iDh8SU4J6FGHIrdVUAVwW7sMknRNTEGhI/XXqLdAbVCNZg
rw46kq0ZhdqLT2nKxhPVQTpOVW/4TIDQKVC/GBQXTOQtzR9KN4smekPKVvigmhtR
/6ksDpG5SlfjC7RV4UJQRQKBgQCubwba0Iolkh/GWvwAup/zqfiTi0Lgtz53kjgw
n8nM8icfyGx7ZoaH+by+FH2yZ42IFpUOQFhdvb5CMNlO1D/SltXVvbWv1+UraDun
IHCU1ECTUN/42JrAy3b5Jh5Ct4Hd+PHebcPCNp9QZrh7Qk7Fo27ajWtH/BIEcnH3
M18jPQKBgQC5M1E9aUYdVRC2j6HyamAm73uLOdltQ7S0pKtypjMYBEUVzbXet2TS
iEkC3ntWGFU9RAOKvwFWvOz3Vuxqgl+H5nkoYH6qhkgAuqOGAHb/DD4VHWzxwHnD
W3SgZHxUij7PPJ8Vslvoov9SZIq5vBZiWvfNKOb4/8KD1IK9dO1aKw==
-----END RSA PRIVATE KEY-----`

var rpPrivK = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwCB4UBzQcnd6GAzPgbt9j2idW23qKZrsvldPNifmOPLfLlMu
sv4EcyJf4L42/aQbTn1rVSu1blGkuCK+oRlKWmZEWh3xv9qrwCwov9Jme/KOE98z
OMB10/xwnYotPadV0de80wGvKT7OlBlGulQRRhhgENNCPSxdUlozrPhrzGstXDr9
zTYQoR3UD/7Ntmew3mnXvKj/8+U48hw913Xn6btBP3Uqg2OurXDGdrWciWgIMDEG
yk65NOc8FOGa4AjYXzyi9TqOIfmysWhzKzU+fLysZQo10DfznnQN3w9+pI+20j2z
B6ggpL75RjZKYgHU49pbvjF/eOSTOg9o5HwX0wIDAQABAoIBAQC4cjOvDYqcadFg
J2RLcvj+5Xs0HFiSqrYfoehc4H8oKxpR+e+6TR1ufxC2zUYzyQmiF8wkTzr19xGA
6XJDbOkx0j5KmbbN7hu2+W4Bgfd7hQgbUct172bvJcnjpJT8PJqqQ0h29oX3veFK
0t1Q4oZW2e3YGUjdO6s39Xro0vGCo7GbUlQSwS93sDQoGhrVmt4hsfjCpXY1+cLA
EY6ZMx1d+R/cMP26AxHrKpyum0VD4cNlsBwbk7gJY5rLCNpMPLNj8ns/0ZjuWRir
1x0diYH1FpoOoOiKq2hU9OdYE6DyNGeTpxSQODKEUwwSEpCJJwKfONp0bnjDqaLK
d+0+LEupAoGBAMWyld7Ej8Au9+w6VgiuHzJHK87PUF73GjAbDmVQJidf49YSBjSS
fOzVAL2Wy4l/fKC73P7mb1U2DyYt3v3ywqPoLsGVg82aN3/btjOpEB/CJKEQct/k
IYOQ82+81MQw06jBv5eB7xjpH9wa6t5lafV32XWLyHgwYUzJe4gm0ltdAoGBAPjJ
UCqB8wHT+hGDwqTM0np+x6uLYSltMPK9HjMr35VvzYpIVlk6qZHxIBHzjg3AsC9V
v65owQB9wt2hsdY3R04XlUuIeOHSZs3OLpacmiqtOsLa53RoZCxSNlCu5NGIAyer
Yxfwx4IdHO73gPDnsYLxrf5bSlL3qb6LMXMPezzvAoGAIl1or9B7LGz9q5J4Ygni
Ylr8wnZHAjrx0mrhlbrY5v9EG3IGohzUmlZsSohr2PrQLyB4ydZEhAthlsFigcIx
E0zI092pi5PDEfafNVut8ddNhrHVRhXhvXz00/d/BJt4L11+cFeluC7N2vTS3tXC
FWk/467oqfu+7hoX3xLgfgECgYB25AW6eqV9zyZnLldbWEKRpXqYISCKopLMvdHr
1GCh0m8gUVdqht04UEnqKkFNkzLfPBRBLfBl4rO4JKiO3ZXm3OBM22ghSuI0If8j
nK0UDfrR2bjYaXbNs3AfeKUC+QPA9meBrmA5bt4/2Om2tpKfKA3lSw0mvxJQa8Zy
3Qgg4wKBgQCkFFf4w0aUUVutEI70B5hebrq37ktj9afpntVoLpb220u7NWsMdwps
2L0EW0Y2pyszOv1TEwgAVsKGcpPzN3FJk861fv64cpn8CIzvBwoA4UWAoOCGkJkD
mtlAVXOCyaIG241q6LvHEcY9S/oc2DtPpnCJB4HiAf3lDLd81Hzkaw==
-----END RSA PRIVATE KEY-----`

var rpPrivK2 = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA1QXXrV7X1b8uFL1PW7+FimlAwxwbEMG5hFru1CN8WsRt8ZVQ
IkXRpiwNNXh1GS0Qmshnv8pKaNCZ5q5wFdUelYspZHVRbIkHiQAaEU5yG9SyavHs
DntUOd50PQ3nC71feW+ff8tvQcJ7+gqf8nZ6UAWpG4bvakPtrJ81h4/Qc23vhtbc
ouP0adgdw6UA0kcdGhTESYMBU0dx/NNysvJhNx36z2UU6kbQ3a2/bINEZAgLfJ7/
Y+/647+tc7bUYdqj3dNkbnk1xiXh5dTLsiow5Xvukpy2uA44M/r2Q5VRfbH2ZrBZ
lgf/XEOZs7zppySgaTWRB5eDTm+YxxyOyykn8wIDAQABAoIBAQCd1/ttInbJkiSi
B3hzImHgIodzSzMe4n0Ffp+zHyw40Y4p0RqUmqly+Pc8pKoX4pWIK3D84vbp3Y/8
J0s0UjucUYZ1Qpz30D1+HU4zfq38w0kFB4eDX40UaCo3R0LpJwREphpIhkRFNMfK
ie7kqTeObfNVS1HBqt3E6B+w+DZcIEI9phmrOcnjEAzPDI4q4sDIUhpHv84tkb/6
lm1RWDlRxgDOGv3knUVXaOvAkTRqdBINKOhaS6dLPpN9FL9aj5UKEklxEtoPSaFP
ib2+RWWe4B+0FPEg0zuSTIH6hhUQK5CBa3CM+0WzfZmsqSpYFbCrmWeOT33tPGy9
NlgVQfwBAoGBAPUdTCFyJPrYdff6VDfxvCDMeLYKJckaa+l3M2Du5BtgFQJM+yPw
5JkNGUyF9MFNWX47cBm0W7pEU8IEiuokhF9XSizX/H8Tz9YyixIVU4krCTtyR6bX
xl36KsB9t8vNtXSN4M8VMMlAZWp5q/n36rzy7jpQKkFfq2todd1yGB3zAoGBAN57
sigfnvxIm41SjAxXnvY8KoP0jCTBxmlxgvLsFhpj9lUIZqQmxbYgeggI9c0MIBNa
/QzmrzLHnSSjtqoUxXy9XY3WKE60uHcvPzePKw5V7EdBWdZLWO+dnKFMziLg1gkx
ccXp0T7VtdQenKRga0PGWw82X/Sr+h90TTmPi04BAoGBAK1Xkb5ZZbOMHylGfAaw
SrX7RCag2IX2zHfn14rmhqShd1oQLM8HDfL643hNh4CoffCagjV7ah85MO6VndPm
DUMLjSZXfHY2AZZeWiFouZHYwIes0uU31U4im9dTUQatLHUH3QM13jGE+/Onpip5
3CTRvA27IZbn3GdyEWCQzmNnAoGAY+laxWgF5rfYmyuB1x0WNvAoC6Aru2oF515h
dyQMfQd9HQyrw3Xh/fsxsiAL+mxCj06iK0QBU6WO7WBT7KdtVKpZtBODgGzqFiPy
mMnDhSmS9SDk7jZiFyFJsKokPEeJ9xDsTfvFyxkAEeU5ZRwjr4kJZZh+mQsORUfe
UkYjQgECgYBEXdVe0uv3V1dUVATfGdBCaZ6BxJVO/VGDKNfoEtFCkR7rVsb+kY6u
DBW8CMBlUhoaSE+/BBEAyyzV++j0nC0cnlU1694HcM0hKx35F4CHPKEAH/ChrTVn
EClJgxaTPuL9ON4s2OaQevT+STkx/dWH/O1FkY5oTAR5JO/wbPZqoQ==
-----END RSA PRIVATE KEY-----`

var IdpPrivK1 = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAwx9oT44DmDRiQJ1K0b9QolEsrQ51hBUDq3oCKTffBikYenSU
QNimVCsVBfNpKhZqpW56hH0mtgLbI7QgZGj9cNBMzSLMolltw0EerF0Ckz0Svvie
1/oFJ1a0Cf4bdKKW6wRzL+aFVvelmNlLoSZXoCpxUPQq7SMLoYEK1c+e3l3H0bfh
6TAVt7APOQEFhXy9MRt83oVSAGW36gdNEksmz1WIT/C1XcHHVwCIJGSdZw5F6Y2g
BjtiLsiFtpKfxQAPwBvDi7uS0PUdN7YQ/G69b0FgoE6qivDTqYfr80Y345Qe/qPG
Dvfne7oA8DIbRV+Kd5s4tFn/cC0Wd+jvrZJ7jwIDAQABAoIBABVYHS/+p/QBXvIU
gre5BtgKqylvGHnPVqxuV0gs/W+OFUhn8kO5r1ArukwBWXKqKxZXpH1Tt2VXoKMi
NBznwzmQ/6W89ceYosImIHXYYsy6dI+BYNbdWaz49g7VxikXFA03WmZWACYIRwwW
UQiayiESI30oiH2SRNZw6D+FS6qlRIR40M6GL3PS+iB6jPslXGencwxhS9jLQopF
47igwzGyO4UZ7eeMG3Nn4IfyCxhHhXqpCE2J0XQIMBtXb9O//tvuAcL3db0jGMTt
v8Q22JxF1X+0/1VfU9JGKfqoqRcOXqix4ztgLXKFel+bj7m64T08MKljD21k3bb5
XrEDfpkCgYEAzg6406B3Peo2kZT5uVh3WzsFuVKAfko0H/vmMXMWM569LtEndjQL
hib4hgkn9co/FBtliC2w03wq/9sKoIPZn73H5EkaAOjyOnF/Z2xvwj/irMJLov9g
i5eHQJljWktzs0aqv2N1r9n6GR3uJDUya07rTpIkWNh3nn5fuO+srjsCgYEA8mo2
CcXGvBeFNIcLR9SFaqNaDPirgoqNmrtxRoaoQvvWVU4rcnXO28E/0vn4l2OUvz5N
IAC+3zEcSTRI//ZGKU6tzJXTbj9BRflLhwhr78D6ArGtyjgWXPBUSp71+qLSsRl5
sHVSsc3acPpMRgGMrlA83OaMnArKGtCcL/qh7r0CgYBPw4EmYpJmBDj1Z963MZia
VyGjGF2nBWBiFSeJcsxgVQ1UhyAocIMZfhJsCDVQvuZmCSjnaxBs/T7D5e2aLw/Z
9yPeqbGIMqQ5nV+9EEu+vO4pA9k1knez8YcoqXe9J0H1XuCPz5dp6A4ZFO3vVCxd
P6J0uruZLMo5LyAsvZJxqwKBgENxKzGS1ZSU0plnjMriJHAjnDUJpeW+mGDZD024
vu1L1TiMc+f3QKLA4/nVU8UCjmqacaiarH+50Q3Ivxp/MMvjONU3RchhTs6h6dJa
lHTyclv3hMtCyW336uuLyBF/5TAiT0m5ilUvWTufV0MOwU3pwtUOS0ZKdin5qcpr
Z0vdAoGABAbnRGHFGBm4jpGhvKT8iXCoRlMvvaYalhpVAYAMehwLR7Chq3O6uTJm
1/iYmfvDJP0ihXWbHJePpTQRnjAu0wOwxeFCS9X1dvFtJEt5DTRbyoZ2hZfwFPxD
GN1apL6Q4DpjL2ktGkYgaKp6HW+5ogfOPPOJOICDScu5Ozl29+o=
-----END RSA PRIVATE KEY-----`

var IdpPrivK2 = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEArdcKj/gAetVyg6Nn2lDim/UJYQsQCav60EVbECm5EVT8Wgnp
zO+GrRyBtxqWUdtGar7d6orLh1RX1ikU7Yx2SA8Xlf+ZDaCELba/85Nb+IppLBdP
ywixgumoto9G9dDGSnPkHAlq5lXXA1eeUS7jiU1lf37lwTZaO0COAuu8Vt9GcwYP
h7SSf4/eXabQGbo/TMUVpXX1w5N1A07Qh5DGr/ZKzEE9/5bJJJRS635OA2T4gIY9
XRWYiTxtiZz6AFCxP92Cjz/sNvSc/Cuvwi15ycS4C35tjM8iT5djsRcR+MJeXyvu
rkaYgMGJTDIWub/A5oavVD3VwusZZNZvpDpDPwIDAQABAoIBAFAfcxDUH3R9+J/P
qsgmy6tSDxaZQLUUfS+NJ+GVOWVRpFXjh80bARm8r9Sy/mGQDS6Z9jJp8lDXgPyG
Rs0OFl40BozuF57+Qq3HM0WSv6sYME1QGUjdIuPRyh8KfoxBw4MBUzvQ42JyYf16
Xs/QKrNX5tYSqNaatI/muw2BlXb8Az/rTdI2i42BLHLDFHo9MS2+bAs+qg/O78pH
/x+H0jNj/5ClL0k7i04WKSCoQNIcFbSOtEtk5ZlaALUWEZavgNn0t8L1ZAd2NK5b
P5VYNwYiyrQhoX6fQkpUcCUwaVq6HnrWtDToBbZnr0pI2dtK+po5CSCLqQ1zr1pd
zRyAlXECgYEA1r5k+n7kGgj5SM1p6vK2KxtnYQ7YYTwTslYxtO+Fy2GhYEqGYTyf
Owe59grnv+RbINERQ3dxsKnC/3DjlS0+uarwfFdVPZYg/7W6FVGfn46FF30/SxqT
AeejFGRgpYGweFHmtCCDRFDCT8BvH0mLRdGI/vMn5p0okZFQaVBG40UCgYEAzzzn
nU1BpkeHMZuzQ+1amReN3Op2afrnr50ho3io8YyfTX2sJmyH0HnlxulJ8cB7bYmN
eqIvXTwLDKPldXaY4NbMXQ6zanbFX/psIPxlRTf2NVqysBBKMAVQg+ZHAUMN0gIs
xw7SCVTaGzSIxmq5FWvHidInf4Dphr2y4SuckrMCgYAE9E+QF+1bTGmz7ElNSlw5
kmBINPd5BtHNg3+SFRSZJJ98gTuocqWZzwvTSV0faD1R/IDRdagB02jUS950Sp7v
2anCtKEa0qPgQmkQpNlx7O/VIuaa7PoHSTjR957jMqLHo9wWu8lLgjF5dY8awa+c
5MCsYR/Cik2tThT02Q1JoQKBgES99CpGlS897Md02Ur/8Zx0prcQAwV2l+G14pGi
FZBCUBlZRYBdYdOyi5imi8OoUIjuJsL2B3YK07N2rkd/doimV5XKqZL4INKMc8+h
SUpjnMTn9/vU+3bgXGvUN9tgTbZKyGWjMeKshciebXw7rHdBkCfUUQvHTC9Iv4xX
dhFnAoGBAL+JIdwLcumdEZJyA3pBu6WfvQJzjDv0HE5N9FsLyTyqc3yacaxH/hv8
nplQvZKsbhmtxbu/MGbJfp1cH0LgO5OamHj5TBEdXWxtZKgE2nmxJz0Fm9L8vZ8b
H5plzde3fZP4YVOa+bK5XuHS5CrwjHoDItfvdPNF8D2rutrl77D9
-----END RSA PRIVATE KEY-----`

var idpPrivK3 = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAz7N55vytQuBV17KHPzd1ILPonOpltFqcMCV+x81NJNcvf2Qc
DDemYK2oObcs8rDuavx3+aSAeBrGXmFIjvVT7YTpEfoCGVf50AJKeyOeuaGefVy1
2GlGUsxKxCWDJaWe6Vc7S+cOyiLHNp/U/La3rSRbJeS6+GLbbVtJZpXsJwIejrK5
2JwSnCTH9aeVUDovJZNfQvPHaKArqermyI7/44o8qfGkImAs4UhLLpcQVyyADaqH
MFKpRTE/cLISCB6Ut9Vb1lyBgk0xlGWLfrXa0erk96NK3tw0thd464qz2qFojNIS
mM1ddG+VSHoZUu7UJzeCUXyw0RkB1PZEXiwz7wIDAQABAoIBACO1JFj8ycC8lqV9
kNjibOWRaIVJmvCVv1Jbr98jwYZ65DSPfm7vRlBKqqg5gKW8m1CTVQD7Mgbz+3SQ
XwwMy0ADYJpxk9jNkiobqrhe2FProDbHMJAjES785kGwfUqEnbxZ/dy/vYAs2Hjg
o5pKw2sl2/G40BgRzs2PKyBS2AWgfSKoh+607mepFNKg0/Hhlvxs3eJv0h9//ez+
himAAg9OY37gkWHS5DlQ7DsQgfFhRUjCFPGcB9Wu3fuMAnOwthtpigDfrp5SqvZI
KrCwJKJkCiqpa7yd9qx4481zTQm/ZyASLiu8CXC9YVgLqhFcxwaOhh5Q0jdcID+Q
5BnGzgECgYEA+MSae2o0a/P1RyBI6IzRo9zDhtTBFA5XZD3tstOoeh0Huhsrp/KR
0L3nxMm9/EouibyW/SwyOVY3doxo9GNN4PWJg3p58bmYkG+ogwZD2VN1s8tAbU6A
YCFtlz9xIx/12Vx5PKL8fr1FQdWeLleI+F3VtrZ8wS083wcbqcySWwECgYEA1b09
+Gv6EGNaGDV3PwbDwAUg9AHTE7TB8QwkB6wkrNN5MuQc40xQpG6xzx/2kf89qEqX
q3gBaaVvFEtH9QCa8sjqix6/1Rs4+V8D8lfjva6FRBtgm2Yaovhrb/ew5npb7KFm
nz/cEUxJ7eXZ2QJJMsGCC0v0OPrlqkCXycp/Pu8CgYBnrfj8is0CWRDW7fu1AEu3
UaEkJrO52ihOHQleSJylGEhKJlzRiGWBbESWXcaSyZAP08vSBIOCJg7Dl81+XYzt
vyfq5jbAqiuNtxuyUAAjKYeawZE+fUM/zW7RZJ2QmBds2f+laAB4CgY9Y/yjL9Rk
Pyd9GR1xnZsLEPlUkXBGAQKBgQCpzI1OrXkbS9JnKRJyn40jHu/u6QQmw5LPTDXT
Yo5APkAqjc3lRNtLxiS7x0i682qoJ5oWPl/g7eww0x13JePyvGqX2vXK9rVsZm9c
NzZVmi+Ey7sTuSmwDmpLqRp//vTIJ/C+0pyhoVmaBN/r5kUAbXpCPzTlj2yktGvh
g11TQQKBgQCtUc8dgxRBEAnlCMjjhVK+8vyl4Tk2dcLL4U7stk+3hstN6bJUFWhl
lsPD7SyWCfa6BdAs9DsTLdUa4EGvfVkRn6oar8OMC+OMhDethRjUmIJV+wWS3ati
I4EPHrPYK3GNb75+G+qH9uJZ1e2FM7CGaDiBHVSthBkCqjEv2e5JGQ==
-----END RSA PRIVATE KEY-----`

var AsPrivK1 = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEApT8lXT9CDRZZkvhZLBD66o7igZf6sj/o0XooaTuy2HuCt6yE
O8jt7nx0XkEFyx4bH4/tZNsKdok7DU75MjqQrdqGwpogvkZ3uUahwE9ZgOj6h4fq
9l1Au8lxvAIp+b2BDRxttbHp9Ls9nK47B3ZuiD02QknUNiPFvf+BWIoC8oe6Abyc
tnV+GTsC/H3jY3BD9ox2XKSE4/xaDMgC+SBU3pqukT35tgOcvcSAMVJJ06B3uyk1
9MzK3MVMm8b4sHFQ76UEpDOtQZrmKR1PH0gVFt93/0FPOH3m4o+9+1OStP51Un4o
H3o80aw5g0EJzDpuv/+Sheec4+0PVTq0K6kjdQIDAQABAoIBACUI4fbkFomYWLr3
rgSSSaoIG/uvdCA+8o8AMc5j8tFR3RoNMBW2Ep1Ah1QYfpPnS2zndO0FqnKmjvWM
nY0EUyijsVAr+uqqIGsFyXqwTf72OC/n5mEQxVFQ9IyOb5npPuMRXAU8upJ+5HAZ
HGGvyVX/Ygm5QjZgDhFnEjYluENii7wbYC7YfHsdkNoUH/10Y+xnfiPpd1X2fD4/
yV7NgrO9uwsLJMHjKELusZwMJ7lQ5JeSHxuTfCKJdyyuaqSij9Q46q/jLuU27Te/
0LuHRl4zggtDRGXfVL4kEjU3B6uXFg2CXFmG+S7mb4zjIhJvTbkj0X7ZrR49kcMB
4btdGAECgYEAze5J0GAPTseiASSN3HI6OIDCvG6iDVpjVszKLsHkfwLDRtr7kuhj
FqAlQjC/dGzvVEWkBC0qZObk2R03E+tI7WdnWsiK01fwbZm7DTtFu9r4kGNvviGK
FNQIuoH8YIn+XMFdjkuzzvw8SorRwWiKiSqDIAyCETdGVlIiuqOhwhECgYEAzWyP
3XL4q+jY8CRCrNczO2lRoTDOH2zpkRxdoF795pGnvXM5M9KCOrcoFjdFPijpgZGR
oIe6hnuipvpIGN13ycIXwa/qIJlALiAq5NSIK/ZKvCB3f6P+gucNs0b7yCceL1qa
A95Qa8eAIoDBAkv1nF+jTZ4DLQAgB6R+g+W/pyUCgYBtjmIyu4gpT0e+9+WI7DRR
Lx9rBCiulfHXkefWbEzVzXB6V7ITfBKLTPPFfQ2+MN46pToXBrhRKg2B/Gr66+fG
dYak46AHw/cjN/Atn+T/hgVLO7uNGWbOoedq4hCUg5WRX0YYl+m3KrYgqi3hiW56
fuV3vW/NHO0Mq3HSfY9nIQKBgByG3++jwK621jF7B5tTAzVT6dcVnPo2OLVDGClm
J6I2RfIEJ0RwDk+zEakMIdyA9/RbT7rYPmngj3TautphHvpwrrXiBQRj48rEAtDm
Rsa8HCLF63JZRsXM6lUkHWDtNb7juRGidM6S1NN1x9fWzpPZoCbuM4izRL9q83rD
k/rVAoGAULKNeMhshK4hghLwYRrvKK+RvqTHzGRRilWqOmVOD5qm9VOpUlLNudao
IdZYlh07pA1L+IXtGdFHL4GlTNa0xXQBsLOTpklqIrTC62ou6026ADM1SC+K/5GE
98StPl4dYJRYvWjfKjfSkqI1J9pV6EPRIHwP+r5gB/EsBqKpmhc=
-----END RSA PRIVATE KEY-----`

var AsPrivK2 = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAzhJ5PP3dfQtpw9p0Kphb30gg9jpgsv425D5pzZaH00zPgYfN
TVZWfrLlTtc/ja8dbHvyDaCyzFD++Vr1vtmSSs9/j8ZhTJrTYHoiHvfG1ulTl1Qd
gwOcrKhpfhhjnCVCPOYjptgac/KPjhT7uiuYwB6axafx+RqPQqwQQhmuuxmTyy69
l/cqezDtYCYUJVA6nV29ZaaF1VjWoE05PK168mcB5quBdE6Vkc4n2k0wxaaTd/s9
LPy6STXtz5IBXH2Gy5RP0TGeXO6iur/ZSM2z/3vQkTMjY/mkDduGioXcB6ieNgVv
3XYbZg4VJEDSuOpRZReKcgLXvwk3CqZZdZRRjQIDAQABAoIBABfdn9jmdc5Tkg4y
sJ12Q72aNucNX8GbG3RXnh1HP7fC/4060xYP17iYs2HsH9oi27+Co0fcwphTERSD
6k4OGJk9asKV8RLUI4La4jS/8XFWWG4AOeLAelassnr+DBs7XW58IMjj4jxnbSTB
XV30Sp6FbNtTVfzJjKnmD4P4QXo9iKOqS1XHGMJ4cb3gizbtvsj4e2U3zVE5N+Um
zvzXcpvyTH1yJrEY2iiorYFUxDzccuWgyQPTOP0Rtn961JlVFVRQmzyf7briVPYJ
s6fjQqt5CH1pTiX1PRzJJx3Bef/6QAsogYDcYH+zZ5xz4ZZeRgPEVXH8tfonUr2E
whOdAykCgYEA/OQcQP67JxjXmRVo0SLSNA6G5fcVwDp875dfu6cnMP1LqNq6NW0w
D5g3prHPYZBz5556zKVitlWsjeq3ABCsxoknwbRBNPEMCWPG+7xi8ie/va8sQWqT
W5vFWlu16difrpCh77sJhOrGV7jwvANmjg2ltiPkKwte9bWIdipKyl8CgYEA0JsE
ODHO2XC+ggRfMItUndccT2GBgi48IEslnDiwVU+JSc9YKh9IxtFa7y43VMkbvncX
qClJ2U78IkPs4OlFbftk0AttqvZfZmWnys8rMZgzNFOhtjgnyBLlRPn5K6z4EjQ+
lEIQzNl0JmMq6x+P5HYxgHHvL5tncUQia9CdA5MCgYBiEvUCH8fk+bVjIPJtaNus
ZJXcSV6eFhCtuj7eP4zratAUw/7DCX1CDv5GH18VrzfD86ocA2es3rz0rLobxFu9
AyPv8z/2kCTi31cj+YNF9jReE7lOBU7wkBCRYk/CSMhkoqKqnhaq/YG+M3Lo90im
fpRtdq3eI6LIF4a8jNpEcQKBgQDGm5A84E8L/qeiqf7m/QCm9nLhsPfYtaRRKrq4
LdDUqFERkPNjxz1G7XQiXGIZuw9LG5/OXuEMoIK1LO6OhAmyWLL20KqtJrxVhVtn
YC7DnSDDJQzFrFlTx4m5TjXJO3lD+7HI/c14+1/2XFw0V2xsG4utusv7C35E/JW5
CHk1OQKBgQCdbhsGrBII6OaAYPKSLSsBQxczbvqE0EP5pH3E5IXKgrQFvkMLBm8D
t+QcA95WPK7J4AEAaU994SeUT1EzHCEU6orxYeFMdSHn3C6/HI9AtnftIPT4lPih
S0ya4kkk1gex6wejZdIAfSEoxNWJd//t9ERfGdGQVOOPsLiiu/W+Ig==
-----END RSA PRIVATE KEY-----`

var AllMasterKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAukTxVg8qpwXebALGCrlyiv8PNNxLo0CEX3N33cR1TNfImItd
5nFwmozLJLM9LpNF711PrkH3EBLJM+qwASlCBayeMiMT8tDmOtv1RqIxyLjEU8M0
RBBedk/TsKQwNmmeU3n5Ap+GRTYoEOwTKNraI8YDfbjb9fNtSICiDzn3UcQj13iL
z5x4MjaewtC6PR1r8uVfLyS4uI+3/qau0zWV+s6b3JdqU2zdHeuaj9XjX7aNV7mv
njYgzk/O7M/p/86RBEOm7pt6JmTGnFu44jBOez6GqF2hZzqR9nM1K4aOedBMHint
Vnhh1oOPG9uRiDnJWvN16PNTfr7XBOUzL03XDQIDAQABAoIBAHOrRlaTun/XlCRs
oICeYnPQKahAuLOa59jCQogzbEgYo5eey+PDRBKlJa0XpQGcMMyQnF9w1TRlBg64
SS2fakFNzTsDL2sUsDOSzcBcDiBAJKKDUJyHsbE0pxdFDi9r7QaXcrtfRqkKFV1U
zB0NsnKOjzJuLiGQVae1QW3FKEDcRBqjEtfGkxpSjGW+zvz23RhMIMfEgdezvtJM
dl1j+u2scy39m1fZuc4bU3fOoppe5LxlqazkVu6WiwbE3dN6QfJhiR8nXMg2RHIh
A5PuI+iqLmmiFW2uu9v4/8ISu/lMlwWU6rTJ0zX9ixO4aonpp6H7KCSmCmXOCM+E
WtMmDGECgYEA9G0O2oYNP+ZN5hbv3w95clbfnIoRZPpBOD8H2jbGlXdowUqQL72m
MfoFpq6VKYbDxPYzPfjZUoeta13bMgu3+4h+LehIR7uTjPvLTFH60Qr4z85RH6AF
SeOIjYBfIaVo/W/7KH0ezVgPZNvq0SwZyv2TwVrkVtMRc6xK1fxC5rkCgYEAwxbo
JS0XwodK5BX4hBh2MxI3r8lx2VZd6DzngwTz2mB3BwjPnqS3HVMOdYr5I4UJqQPn
HTtRrEIIdynh81n5H5Qkw0DVzGVmpJSItsJBzr46grY9p0OZfYZyYmeV+N28JNEY
QPqpFZSI2oyCn7km/2eqj2YYM75zjoBBrhP7SPUCgYEA2kzcy0aWZs+mGy25JpuH
eBsms4SMbIcl4LpKpRXu3mc7ZAbYKAtVd6U5jti119TI3AyXT24FirQqqo20y0m0
FC6foximFYruCSiJNayyOil2dwJpabldf9R7jQVt8Xrt/gwZYNv+up8/gHD5k7+z
eZxobnRjIzh3ibwDSoJ2reECgYEAgKgWqI24YZ1/kjO7FMJdEQkumEstPbtrasDf
nNQjTRzY4la5NVJDQJ+JpZLlAru1xzS/sdNw5T0XAB8q16W6WU0FgY68cHNe4aLj
FkO9ym5Bf/pXZnt6OgH0ZVkS2nDApzcN26xy3bx7FEYdzt/4C+9919vokhdDdfK3
XennihECgYEAo7gsAeulbJbeLDv4KntoRwl49n1YkO5yKza/icSASo/Qpk9n8Tr7
0Tr4yud9RbBwpnMsuj0FBwS6xrED+sqRlhl4Hg6rdPcj8Jl9+WndcwLjDLdo0Xrc
dEtsboiI40w8gQBC7GvFQ9ihmYQDOyhOlPLCmc4w2Yg2nkyfctJ9kcA=
-----END RSA PRIVATE KEY-----`

var idpPrivK4 = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEApbxaA5aKnkpnV7+dMW5x7iEINouvjhQ8gl6+8A6ApiVbYIzJ
CCaexU9mn7jDP634SyjFNSxzhjklEm7qFPaHOk1FfX6tk5i5uGWifRQHueXhXjR8
HSBkjQAoZ0eqBqTsxsSpASsT4qoBKtsIVN7XHdh9Mqz+XAkq4T6vtdaocduarNG6
ALZFkX+pAgkCj4hIhRmHjlyYIh1yOZw1KM3THkM9noP2AYEH2MBHCzuu+bifCwur
OBq+ZKAdfroCG4rPGfOXuDQK8BHpru1lg0jdAmbbqMyGpAsF+WjW4V2rcTMFZOoY
FYE5m2ssxC4O9h3f/H2gBtjjWzYv6bRC6ZdP2wIDAQABAoIBAHnL21K7tQ7ymtOP
i1OiWLOpLsH3EYKWOImOWz9LSRvQZECl9a65wwA5g69pNoN7s/Z39cVH73X6VNYh
EIFrUqFz29eH2sOW/xUWC71jlPH2kBKM+5DkF0DPluGfdsH/PcotCA5FvA1c5hK6
eHr2cJwMVqWBIEQ+sHZrfPFi2NMiYl7RB0gxwt+CY6ezDUY7TOqdx/3UIDXQa9nA
PQ2GvV8cfKQfLl3rfsoF2ObO+fs6Kmsp3aYJxRVu8aaD+UKS+ljPsD+/IYGa/8Z1
ixXVSHscDrIfRR7LmYDXtOal1pwo0gdErNAMivPXwY1XZ03iHTm7pihhTFd86jDs
9nLRjUECgYEA2jC3sPHDLEeXn8tVMj8B5VBSdFi+H5tyzJbR5ef0QAl7vslbi5sq
MGXtFhRMDjeTwLBAa5YMnYhqHV53GB2zK95+FJeQy1rfSQsxZwIFL0L4wySP19EI
aXaUpPDgdRKS1SC05auiQp5PE2AxmIHfBc77HDE8iU5SiS2TOKFf+isCgYEAwnSp
0Pa/lkHqfEqe6I56Zoi28eUGwjVWc/USdNsM8sfCja49UCryHHHljm1PfgFX4IFy
Kzupm7l1cHQsFI3o1VYWxo/DvXKJ6NjGZLvhbWMLsAaMptcWrkMZ/NDDThuyg9hg
QE3cvNYUk2eoy4PKcyqPeXggYd7Ue54EO+scmRECgYBya5no8N+pGOIqqjbDYsdb
ugODgAY0DRDmuTDZoAo2isKaCn43d+dn+guayIoZ6otRQRyHTujOs/rx69gIjYqo
NsVnhxQnkEAHzhbaLfUKE9TggQvt4XDH3aeV17vdqR/XJI+44Yj15o8RWiCoGXMb
WK/W2PsmBizCQ2QxDm+GgQKBgEicd6zn9rKM+ppe4ufEDECtXGMHOnbao+W45aNt
CHC/1w5AufRtlOq6PRXqC3zp036p14/9P2A+6HONbchfFUpUUzziAh2D36trBuom
ng7SpVKdn3fNaVK5C8Mz0TohbY9+BLL+YCbDafuBAa69D6PhiKG7EZx6MK3YW4xk
RtGBAoGBAMMb1JaB7GF7Yia/8K98PCUeIYOKiPL2XVo0Vpf8z59cytvo/l3VAXXe
tMtaNcpcjJxooc3p9fVlZrU39Zs2fwL53qwlrsIhDFcGWwSCZB6Bcec+LD+MSTr+
rdDFuWTHd+msuASyCbJ2vX3SwnL7g/vapR24umMI6lrJqFGu2aSF
-----END RSA PRIVATE KEY-----`

var idpPrivK5 = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAu9+CK/vznpXtAUC0QhuJgYKCfMMBiIgVcp2A+e+SsKvv6ESQ
72R8K6nQAhH2MGtnj3ScLI0tMwCtgotWCEGiyUXKXLVTiqAqtwflCUVuxCDVuvOm
3GQCxvwzE34jEgbGZ33G3tV7uKTtifhoJzVYD+WkZVslBhaBgQCUewCX4zkCCTYC
5VEhkr7K8HGEr6n1eBOO5VORCkrHKYoZK7euNjyWvWYyVN07F8K0RhgIF9Xsa6Ti
u1Yf8zuyJ/awR6U4Nw+oTkvRpx64+caBNYgR4n8peg9ZJeTAwV49o1ymx34pPjHU
gSdpyhZX4i3z9ji+o7KbNkA/O0l+3doMuH1exwIDAQABAoIBAGyLLcILxy0QoeXf
ZEXtcvyIUquSXwhq1zlpFmNQrwezzt/6/WHSRItViQApMHu5EhQn4zM6PasB8T1D
E2mhwlNXJxt5B9NHxmYJAaLhoqVd8x4YN4eNoK0meLwCXHDFyUtxt7x2ywxa/YKB
Kmu8viwxGVIV3sYtqpTFqQOHzDlScRLv+qsgXw4OimjOTtvz1Gas7MnGnSweTYOo
6CAs3daORZ+tWwvwnAwjVIJItYC3T1hOOLJHkiB47EcKONjmt5jjy1Qvx7QRjK2c
/WOjGsO2tTAqfSW0jiEnybHTBqEqEKyufkDHKaChtur8sq5FFKdRaxnqKAfnQE3e
NaIMosECgYEA3lHLQT3kNUJk7UzxaBTlfmyaUG7pI4teFiPOCm6NQwQyiatPWSkt
Bo6uZ54r0Om2N6gPGl5rNC9M9njq0vPyOAKg2HT1E4LeoGPonEu65wiRdclLbtkW
NId7DYEtTCBI/imfncwSnMtLPTbnlZCCjQZljcuCk6A4hJ0SnxOS7FMCgYEA2FXH
zfcTt3rvTwMfrK44lxotOftlWbwu4nIjZQ57E+2JIV4eBg4hoOV3YoYVDb0nqBwk
5MGFJqrQoL9/1KbV/hkYZ4NDaCtT1U4iZ168rmbtQ1tO1QrA5D5IXjS0khC6syzO
716qZI1Y+xqQbnEEjEpzYnmwFL1PH6XFnqV+1T0CgYEArkV1y92lPy6diPrwnYML
5s9hI73dWXSNO1Oz1q+UYj0vFIXKPH0vg11jT2xIsooRwY0m0afD53NQpEBi6xw4
+jjtNuBvoGzM8POASsx+ZU5tH+S8EddwNZsiFZL2HB+OuFWOfpaS3H/rqb+ZR7+w
5rVl9AHciLZmt2WdTD9+w2sCgYEAqXwK3UIFIGofsjcwSYj0rOzFIffinzrfQGlL
cZC2vBYMqSejPfs0PWmI7pc9R1Y6C2qBPPaf6ntIl6dv7poGbNwcUnx0AthvBV4B
dhqyl6/rkimmySFznV1uNN/117lji5w/QylXNQ/H9nIJVX0VoxNw8mWDnbvykUi+
Wlwt0cECgYA9jlWS+rLppVflVTA4CAMNfXt8qyjDc3Wl3HUB0rgApRXbkMQb875U
6czVLqkWbRiIabfPXYrJRaTjG8thQuquzpaS6+U/ed6t0vUJuxTXkmj1HS3a/dJT
VTVj/BlsFEVvTc0wuiA3mwlgNirRI1UH0GYWlk22hUF3MpMF46SAVQ==
-----END RSA PRIVATE KEY-----`

var AccessorPubKey1 = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7BjIuleY9/5ObFl0w+U2
fID4cC8v3yIaOjsImXYNon04TZ6lHs8gNvrR1Q0MRtGTugL8XJPj3tw1AbHj01L8
W0HwKpFQxhwvGzi0Sesb9Lhn9aA4MCmfMG7PwLGzgdeHR7TVl7VhKx7gedyYIdju
EFzAtsJYO1plhUfFv6gdg/05VOjFTtVdWtwKgjUesmuv1ieZDj64krDS84Hka0gM
jNKm4+mX8HGUPEkHUziyBpD3MwAzyA+I+Z90khDBox/+p+DmlXuzMNTHKE6bwesD
9ro1+LVKqjR/GjSZDoxL13c+Va2a9Dvd2zUoSVcDwNJzSJtBrxMT/yoNhlUjqlU0
YQIDAQAB
-----END PUBLIC KEY-----`

var AccessorPubKey2 = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhdKdvawPO8XXroiAGkxF
fLRCqvk4X2iAMStq1ADjmPPWhKgF/ssU9LBdHKHPPX1+NMOX29gOL3ZCxfZamKO6
AbODt1e0bVfblWWMq5uMwzNrFo4nKas74SLJwiMg0vtn1NnHU4QTTrMYmGqRf2WZ
IN9Iro4LytUTLEBCpimWM2hodO8I60bANAO0gI96BzAWMleoioOzWlq6JKkiDsj7
8EjCI/bY1T/v4F7rg2FxrIH/BH4TUDy88pIvAYy4nNEyGyr8KzMm1cKxOgnJI8On
wT8HrAJQ58T3HCCiCrKAohkYBWITPk3cmqGfOKrqZ2DI+a6URofMVvQFlwfYvqU6
5QIDAQAB
-----END PUBLIC KEY-----`
