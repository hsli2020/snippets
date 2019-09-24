<?php

$doc = new DOMDocument();

$xmlStream = <<<MyFeed
<?xml version="1.0"?>
<feed xmlns="http://www.w3.org/2005/Atom" xml:lang="en">
  <title>Using XPath with PHP</title>
  <author>
    <name>Tracy Bost</name>
  </author>
  <subtitle type="html">
    Let XPath do the hard work for you when working with XML</subtitle>
  <link rel="self" type="text/html" hreflang="en" href="http://www.ibm.com/developerworks/"/>
  <updated>15 Aug 2011 22:51:48 +0000</updated>
  <entry>
    <title>SimpleXML and XPath </title>
    <summary>If you are using SimpleXML to parse XML or
         RSS feeds, XPath is great to use!</summary>
    <link rel="self" type="text/html" hreflang="en" href=""/>
    <published>21 Apr 2011 04:00:00 +0000</published>
    <updated>21 Apr 2011 04:00:00 +0000</updated>
  </entry>
  <entry>
    <title>DOMXPath</title>
    <summary>If you are using DOM for traversal XML documents, 
        give DOMXPath a try! </summary>
    <link rel="self" type="text/html" hreflang="en" href=""/>
    <id>tag:developerWorks.dw,19 Apr 2011 04:00:00 +0000</id>
    <published>12 Aug 2011 04:00:00 +0000</published>
    <updated>12 Aug 2011 04:00:00 +0000</updated>
  </entry>
  <entry>
    <title>XMLReader with XPath</title>
    <summary>For complex XML document reading and writing, 
        using XPath with XReader can ease your burden!</summary>
    <link rel="self" type="text/html" hreflang="en" href=""/>
    <id>tag:developerWorks.dw,19 Apr 2011 04:00:00 +0000</id>
    <published>08 Aug 2011 04:00:00 +0000</published>
    <updated>08 Aug 2011 04:00:00 +0000</updated>
  </entry>
</feed>
MyFeed;

$xsldoc = <<<XSL
<?xml version='1.0'?>
<xsl:stylesheet version="1.0"
                xmlns:atom="http://www.w3.org/2005/Atom"
                xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:dc="http://purl.org/dc/elements/1.1/">
  <xsl:template match="/">
    <html>
      <head><title><xsl:value-of select="//atom:title"/></title></head>
      <table>
        <tr><td><xsl:value-of select="//atom:title"/></td></tr>
        <tr><td><i>"<xsl:value-of select="//atom:subtitle" />",</i></td></tr>
        <tr><td>by <xsl:value-of select="//atom:author"/></td></tr>
        <xsl:for-each select="//atom:feed/entry">
          <table border="1" >
            <tr>
              <td>Title</td><td><xsl:value-of select="//atom:title"/></td>
            </tr>
            <tr>
              <td>Summary</td><td><xsl:value-of select="//atom:summary"/></td>
            </tr>
            <tr>
            </tr>
          </table><br/>
        </xsl:for-each>
      </table>
    </html>
  </xsl:template>
</xsl:stylesheet>
XSL;

$doc->loadXML($xmlStream);
//$xpath = new DOMXpath($doc);

$xsl = new DOMDocument();
$xsl->loadXML($xsldoc, LIBXML_NOCDATA);

$xslt = new XSLTProcessor();
$xslt->importStylesheet($xsl);

print $xslt->transformToXML($doc);
