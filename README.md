# doffin_extractor
<h1>Extracts Doffin announcements for search</h1>

<h2>Doffin is the Norwegian Database for public procurements (doffin.no)</h2>

IMPORTANT: the doffin data is downloaded into a file. Download will fail if the directory executed from is not writable.<br>

Usage:<br>
  doffin_extractor [command]<br>

Available Commands:<br>
  description&emsp;Search by description<br>
  heading&emsp;Search by heading<br>
  help&emsp;Help about any command<br>
  id&emsp;Search by ID<br>
  load&emsp;Load data from Doffin<br>
  recent&emsp;Search for recent postings<br><br>

Flags:
  -h&emsp;, --help&emsp;help for doffin_extractor<br>
  -t,&emsp; --toggle&emsp;Help message for toggle<br>
  -v,&emsp; --version&emsp;version for doffin_extractor<br><br>

Use "doffin_extractor [command] --help" for more information about a command.

To build:<br>
Install golang: https://go.dev/doc/install<br>
Clone<br>
go build .<br>
go install .<br>


