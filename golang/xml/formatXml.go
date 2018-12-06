import (
    "bytes"
    "encoding/xml"
    "io"
)

func formatXML(data []byte) ([]byte, error) {
    b := &bytes.Buffer{}
    decoder := xml.NewDecoder(bytes.NewReader(data))
    encoder := xml.NewEncoder(b)
    encoder.Indent("", "  ")
    for {
        token, err := decoder.Token()
        if err == io.EOF {
            encoder.Flush()
            return b.Bytes(), nil
        }
        if err != nil {
            return nil, err
        }
        err = encoder.EncodeToken(token)
        if err != nil {
            return nil, err
        }
    }
}

func Prettify(raw string, indent string) (pretty string, e error) {
    r := strings.NewReader(raw)
    z := html.NewTokenizer(r)
    pretty = ""
    depth := 0
    prevToken := html.CommentToken
    for {
        tt := z.Next()
        tokenString := string(z.Raw())

        // strip away newlines
        if tt == html.TextToken {
            stripped := strings.Trim(tokenString, "\n")
            if len(stripped) == 0 {
                continue
            }
        }

        if tt == html.EndTagToken {
            depth -= 1
        }

        if tt != html.TextToken {
            if prevToken != html.TextToken {
                pretty += "\n"
                for i := 0; i < depth; i++ {
                    pretty += indent
                }
            }
        }

        pretty += tokenString

        // last token
        if tt == html.ErrorToken {
            break
        } else if tt == html.StartTagToken {
            depth += 1
        }
        prevToken = tt
    }
    return strings.Trim(pretty, "\n"), nil
}

/*
html := `<!DOCTYPE html><html><head>
<title>Website Title</title>
</head><body>
<div class="random-class">
<h1>I like pie</h1><p>It's true!</p></div>
</body></html>`
pretty, _ := Prettify(html, "    ")
fmt.Println(pretty)
*/


func prettyPrint(b *bytes.Buffer, n *html.Node, depth int) {
    switch n.Type {
    case html.DocumentNode:
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            prettyPrint(b, c, depth)
        }

    case html.ElementNode:
        justRender := false
        switch {
        case n.FirstChild == nil:
            justRender = true
        case n.Data == "pre" || n.Data == "textarea":
            justRender = true
        case n.Data == "script" || n.Data == "style":
            break
        case n.FirstChild == n.LastChild && n.FirstChild.Type == html.TextNode:
            if !isInline(n) {
                c := n.FirstChild
                c.Data = strings.Trim(c.Data, " \t\n\r")
            }
            justRender = true
        case isInline(n) && contentIsInline(n):
            justRender = true
        }
        if justRender {
            indent(b, depth)
            html.Render(b, n)
            b.WriteByte('\n')
            return
        }
        indent(b, depth)
        fmt.Fprintln(b, html.Token{
            Type: html.StartTagToken,
            Data: n.Data,
            Attr: n.Attr,
        })
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            if n.Data == "script" || n.Data == "style" && c.Type == html.TextNode {
                prettyPrintScript(b, c.Data, depth+1)
            } else {
                prettyPrint(b, c, depth+1)
            }
        }
        indent(b, depth)
        fmt.Fprintln(b, html.Token{
            Type: html.EndTagToken,
            Data: n.Data,
        })

    case html.TextNode:
        n.Data = strings.Trim(n.Data, " \t\n\r")
        if n.Data == "" {
            return
        }
        indent(b, depth)
        html.Render(b, n)
        b.WriteByte('\n')

    default:
        indent(b, depth)
        html.Render(b, n)
        b.WriteByte('\n')
    }
}

func isInline(n *html.Node) bool {
    switch n.Type {
    case html.TextNode, html.CommentNode:
        return true
    case html.ElementNode:
        switch n.Data {
        case "b", "big", "i", "small", "tt", "abbr", "acronym", "cite", "dfn", "em", "kbd", "strong", "samp", "var", "a", "bdo", "img", "map", "object", "q", "span", "sub", "sup", "button", "input", "label", "select", "textarea":
            return true
        default:
            return false
        }
    default:
        return false
    }
}

func contentIsInline(n *html.Node) bool {
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        if !isInline(c) || !contentIsInline(c) {
            return false
        }
    }
    return true
}

func indent(b *bytes.Buffer, depth int) {
    depth *= 2
    for i := 0; i < depth; i++ {
        b.WriteByte(' ')
    }
}

func prettyPrintScript(b *bytes.Buffer, s string, depth int) {
    for _, line := range strings.Split(s, "\n") {
        line = strings.TrimSpace(line)
        if line == "" {
            continue
        }
        depthChange := 0
        for _, c := range line {
            switch c {
            case '(', '[', '{':
                depthChange++
            case ')', ']', '}':
                depthChange--
            }
        }
        switch line[0] {
        case '.':
            indent(b, depth+1)
        case ')', ']', '}':
            indent(b, depth-1)
        default:
            indent(b, depth)
        }
        depth += depthChange
        fmt.Fprintln(b, line)
    }
}
