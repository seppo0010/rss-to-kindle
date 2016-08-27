package content

//ArticleTmpl ...
const ArticleTmpl = `<html>
  <head>
    <meta content="text/html; charset=utf-8" http-equiv="Content-Type"/>
    <title>%s</title>
  </head>
  <body>
    <div class="header">
      <h1>%s</h1>
    </div>
    %s
  </body>
</html>`

//ImageTagTmpl ...
const ImageTagTmpl = `<p><img src="%s.jpg" middle="true" style="margin-bottom: 20px"></p>`

//ContentPageTmpl ...
const ContentPageTmpl = `<html>
  <head>
    <meta content="text/html; charset=utf-8" http-equiv="Content-Type"/>
    <title>Table of Contents</title>
  </head>
  <body>
    <h1>Contents</h1>
    %s
  </body>
</html>`

//ContentSectionTmpl ...
const ContentSectionTmpl = `<h4>%s</h4><ul>%s</ul>`

//ContentArticleTmpl ...
const ContentArticleTmpl = `<li><a href="%s">%s</a></li>`

//ManifestItemTmpl ...
const ManifestItemTmpl = `<item href="%s" media-type="%s" id="%s"/>`

//NavHeaderTmpl ...
const NavHeaderTmpl = `<head>
  <meta content="%s" name="dtb:uid"/>
  <meta content="2" name="dtb:depth"/>
  <meta content="0" name="dtb:totalPageCount"/>
  <meta content="0" name="dtb:maxPageNumber"/>
</head>
<docTitle>
  <text>%s/text>
</docTitle>
<docAuthor>
  <text>$s</text>
</docAuthor>`

//NavSectionTmpl ...
const NavSectionTmpl = `<navPoint playOrder="%d" class="section" id="%s">
  <navLabel>
    <text>%s</text>
  </navLabel>
  <content src="%s"/>
  %s
</navPoint>`

//NavArticleTmpl ...
const NavArticleTmpl = `<navPoint playOrder="%d" class="article" id="%s">
  <navLabel>
    <text>%s</text>
  </navLabel>
  <content src="%s"/>
  <mbp:meta name="description">%s</mbp:meta>
  <mbp:meta name="author">%s</mbp:meta>
</navPoint>
`

//NavMapTmpl ...
const NavMapTmpl = `<navMap>
  <navPoint playOrder="0" class="periodical" id="periodical">
    <navLabel>
      <text>Table of Contents</text>
    </navLabel>
    <content src="contents.html"/>
    %s
  </navPoint>
</navMap>`

//NavMainTmpl ...
const NavMainTmpl = `<?xml version='1.0' encoding='utf-8'?>
<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN" "http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
<ncx xmlns:mbp="http://mobipocket.com/ns/mbp" xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1" xml:lang="en-GB">
  %s
  %s
</ncx>`

//OpfTmpl ...
const OpfTmpl = `<?xml version='1.0' encoding='utf-8'?>
<package xmlns="http://www.idpf.org/2007/opf" version="2.0" unique-identifier="%s">
	<metadata>
		<dc-metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
			<dc:title>%s</dc:title>
			<dc:language>%s</dc:language>
			<meta content="cover-image" name="cover"/>
			<dc:creator>%s</dc:creator>
			<dc:publisher>%s</dc:publisher>
			<dc:subject>News</dc:subject>
			<dc:date>%s</dc:date>
			<dc:description>%s</dc:description>
		</dc-metadata>
		<x-metadata>
			<output content-type="application/x-mobipocket-subscription-magazine" encoding="utf-8"/>
		</x-metadata>
	</metadata>
	<manifest>
	%s
	</manifest>
	<spine toc="nav-contents">
	%s
	</spine>
	<guide>
		<reference href="contents.html" type="toc" title="Table of Contents"/>
	</guide>
</package>`

//SpineTmpl ...
const SpineTmpl = `<itemref idref="%s"/>`

//ImageTagRegex ...
const ImageTagRegex = `<img\s[^>]*?src\s*=\s*['\"]([^'\"]*?)['\"][^>]*?>`
