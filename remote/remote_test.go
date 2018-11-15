package remote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var content = `
<table class="codetable">
<thead>
<tr class="first">
  <th>File name</th>
  <th>Kind</th>
  <th>OS</th>
  <th>Arch</th>
  <th>Size</th>
  
  <th>SHA256 Checksum</th>
</tr>
</thead>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.src.tar.gz">go1.7.3.src.tar.gz</a></td>
  <td>Source</td>
  <td></td>
  <td></td>
  <td>14MB</td>
  <td><tt>79430a0027a09b0b3ad57e214c4c1acfdd7af290961dd08d322818895af1ef44</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.darwin-amd64.tar.gz">go1.7.3.darwin-amd64.tar.gz</a></td>
  <td>Archive</td>
  <td>OS X</td>
  <td>64-bit</td>
  <td>78MB</td>
  <td><tt>2ef310fa48b43dfed7b4ae063b5facba130ed0db95745c538dfc3e30e7c0de04</tt></td>
</tr>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.darwin-amd64.pkg">go1.7.3.darwin-amd64.pkg</a></td>
  <td>Installer</td>
  <td>OS X</td>
  <td>64-bit</td>
  <td>78MB</td>
  <td><tt>c2b0e222ab32c92283fbebea61d54c1fbd015d94654384e0fc40162a68898c22</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.freebsd-386.tar.gz">go1.7.3.freebsd-386.tar.gz</a></td>
  <td>Archive</td>
  <td>FreeBSD</td>
  <td>32-bit</td>
  <td>69MB</td>
  <td><tt>e3ac58b1ea8272570adb646bcf4f313d52afe453c83f155ef3f931f472261f0e</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.freebsd-amd64.tar.gz">go1.7.3.freebsd-amd64.tar.gz</a></td>
  <td>Archive</td>
  <td>FreeBSD</td>
  <td>64-bit</td>
  <td>79MB</td>
  <td><tt>78e8987603ab379c9aa1707e027e46978a26f71caf5c0df4cf3a4627570efff5</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.linux-386.tar.gz">go1.7.3.linux-386.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>32-bit</td>
  <td>69MB</td>
  <td><tt>d39d562c3247b11ae659afe1e131a3287c60b7de207ca5f25684c26f1c1dff5c</tt></td>
</tr>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz">go1.7.3.linux-amd64.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>64-bit</td>
  <td>79MB</td>
  <td><tt>508028aac0654e993564b6e2014bf2d4a9751e3b286661b0b0040046cf18028e</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.linux-armv6l.tar.gz">go1.7.3.linux-armv6l.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>ARMv6</td>
  <td>66MB</td>
  <td><tt>d02912d121e1455e775a5aa4ecdb2a04f8483ba846e6d2341e1f35b8e507d7b5</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.linux-s390x.tar.gz">go1.7.3.linux-s390x.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>s390x</td>
  <td>66MB</td>
  <td><tt>cadbf9cab94c91b4e8d37884cbe4dd237f983b4c92238c0e93628c166440fb50</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.windows-386.zip">go1.7.3.windows-386.zip</a></td>
  <td>Archive</td>
  <td>Windows</td>
  <td>32-bit</td>
  <td>74MB</td>
  <td><tt>d0ac2d3aaa20452d0f09112f034cca1c5b8560452a45e10523af7f0a1089c792</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.windows-386.msi">go1.7.3.windows-386.msi</a></td>
  <td>Installer</td>
  <td>Windows</td>
  <td>32-bit</td>
  <td>62MB</td>
  <td><tt>dd2ae25fc099003f5207c6125ceb4cd9444a866d4d28084a653e178514d41727</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.windows-amd64.zip">go1.7.3.windows-amd64.zip</a></td>
  <td>Archive</td>
  <td>Windows</td>
  <td>64-bit</td>
  <td>85MB</td>
  <td><tt>9fe41313b97e2a6a703f5ae22938c7d9ac4336a128b522376c224ba97e8c7f01</tt></td>
</tr>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.7.3.windows-amd64.msi">go1.7.3.windows-amd64.msi</a></td>
  <td>Installer</td>
  <td>Windows</td>
  <td>64-bit</td>
  <td>72MB</td>
  <td><tt>1e73eb0c36af1b714389328a39de28254aca35956933ec84ca2d93175471f41a</tt></td>
</tr>

</table>

<table class="codetable">
<thead>
<tr class="first">
  <th>File name</th>
  <th>Kind</th>
  <th>OS</th>
  <th>Arch</th>
  <th>Size</th>
  
  <th>SHA1 Checksum</th>
</tr>
</thead>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.2.2.windows-amd64.msi">go1.2.2.windows-amd64.msi</a></td>
  <td>Installer</td>
  <td>Windows</td>
  <td>64-bit</td>
  <td></td>
  <td><tt>c8f5629bc8d91b161840b4a05a3043c6e5fa310b</tt></td>
</tr>

</table>
<table class="codetable">
<thead>
<tr class="first">
  <th>File name</th>
  <th>Kind</th>
  <th>OS</th>
  <th>Arch</th>
  <th>Size</th>
  
  <th>SHA256 Checksum</th>
</tr>
</thead>

  
<tbody><tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.src.tar.gz">go1.9beta2.src.tar.gz</a></td>
  <td>Source</td>
  <td></td>
  <td></td>
  <td>16MB</td>
  <td><tt>4ca11b29e9c3b2ef1db837a80bc3a54a6ba392dc3f7447cb99972f9c96daa8c3</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.darwin-amd64.tar.gz">go1.9beta2.darwin-amd64.tar.gz</a></td>
  <td>Archive</td>
  <td>OS X</td>
  <td>64-bit</td>
  <td>90MB</td>
  <td><tt>07e25fa6e290e46bee8ee2d836d79c3f5a35917b013ecf686d7e761631524de2</tt></td>
</tr>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.darwin-amd64.pkg">go1.9beta2.darwin-amd64.pkg</a></td>
  <td>Installer</td>
  <td>OS X</td>
  <td>64-bit</td>
  <td>90MB</td>
  <td><tt>af7a9e81bff46d9a5e84362d285fe0827a98e5af56ef57a9aa7088693cd24489</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.freebsd-386.tar.gz">go1.9beta2.freebsd-386.tar.gz</a></td>
  <td>Archive</td>
  <td>FreeBSD</td>
  <td>32-bit</td>
  <td>79MB</td>
  <td><tt>27283b4b96fe9f566d53e26b4ce7db10cd4d2d5bfdd37914c19fe44e4c39bdb3</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.freebsd-amd64.tar.gz">go1.9beta2.freebsd-amd64.tar.gz</a></td>
  <td>Archive</td>
  <td>FreeBSD</td>
  <td>64-bit</td>
  <td>90MB</td>
  <td><tt>86cdd9759802d1b985978b4158f6e91cfae8b43a52fc99219dfc6402ee356036</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.linux-386.tar.gz">go1.9beta2.linux-386.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>32-bit</td>
  <td>79MB</td>
  <td><tt>b06f1ef695560d687c94ebd83a0c2f6d1adfa37121ddd14d361816eaa2f00d44</tt></td>
</tr>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.linux-amd64.tar.gz">go1.9beta2.linux-amd64.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>64-bit</td>
  <td>90MB</td>
  <td><tt>023f778f063d2234e7c95f572a92298b307807693f7e045a88c90ecd7a08f29d</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.linux-arm64.tar.gz">go1.9beta2.linux-arm64.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>arm64</td>
  <td>77MB</td>
  <td><tt>4e60b704f04441ad97b5a7c660a680225abd59b33b9044731066f2f91c18ddba</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.linux-armv6l.tar.gz">go1.9beta2.linux-armv6l.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>ARMv6</td>
  <td>78MB</td>
  <td><tt>4be3d357c7a6ccd51c2fb89efbe32da3fa1d209f105f146429a7ea2c015e64dc</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.linux-ppc64le.tar.gz">go1.9beta2.linux-ppc64le.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>ppc64le</td>
  <td>76MB</td>
  <td><tt>52aad09c7b858a02d3de72ea971b656da698ef8be2647277d751b2210d9d1419</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.linux-s390x.tar.gz">go1.9beta2.linux-s390x.tar.gz</a></td>
  <td>Archive</td>
  <td>Linux</td>
  <td>s390x</td>
  <td>76MB</td>
  <td><tt>c8166dfb2af4290acffb9382ca3a5723e8b400c718cb634614c3248740c2067f</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.windows-386.zip">go1.9beta2.windows-386.zip</a></td>
  <td>Archive</td>
  <td>Windows</td>
  <td>32-bit</td>
  <td>84MB</td>
  <td><tt>d998f29e2c8a774b2ddde13a282c5b146b365b98022a1ca7a16068577036d2ba</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.windows-386.msi">go1.9beta2.windows-386.msi</a></td>
  <td>Installer</td>
  <td>Windows</td>
  <td>32-bit</td>
  <td>71MB</td>
  <td><tt>527b332aa4507ffff00a4cc5689038d1ceeb22cb8324c9332640aa3f3beb58fa</tt></td>
</tr>

<tr>
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.windows-amd64.zip">go1.9beta2.windows-amd64.zip</a></td>
  <td>Archive</td>
  <td>Windows</td>
  <td>64-bit</td>
  <td>97MB</td>
  <td><tt>692460aacdbfa2c36765955ede95b1c6c8b3b7afdbe8f9149947ac81e4378f2e</tt></td>
</tr>

<tr class="highlight">
  <td class="filename"><a class="download" href="https://storage.googleapis.com/golang/go1.9beta2.windows-amd64.msi">go1.9beta2.windows-amd64.msi</a></td>
  <td>Installer</td>
  <td>Windows</td>
  <td>64-bit</td>
  <td>82MB</td>
  <td><tt>9dc9e29d1613ba0979df04df6928b7f962f1a61c17408799dad3a4e5652fe3d9</tt></td>
</tr>


</tbody></table>
`

func TestAvailableVersions(t *testing.T) {
	remote := Remote{content, []Version{}}
	remote.Update()
	assert.Equal(t, remote.Versions[0].Release, "1.7.3")
	assert.Equal(t, remote.Versions[0].arch, "amd64")
	assert.Equal(t, remote.Versions[0].os, "linux")
}
