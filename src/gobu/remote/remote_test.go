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

</table>`

func TestAvailableVersions(t *testing.T) {
	remote := Remote{content, []Version{}}
	remote.Update()
	assert.Equal(t, remote.versions[0].version, "1.7.3")
	assert.Equal(t, remote.versions[0].arch, "amd64")
	assert.Equal(t, remote.versions[0].os, "linux")
}
