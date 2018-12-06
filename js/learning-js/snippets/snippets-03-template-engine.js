// JavaScript template engine in just 20 lines
// http://krasimirtsonev.com/blog/article/Javascript-template-engine-in-just-20-line
var TemplateEngine = function(html, options) {
    var re = /<%([^%>]+)?%>/g,
        reExp = /(^( )?(if|for|else|switch|case|break|{|}))(.*)?/g,
        code = 'var r=[];\n', cursor = 0, match;

    var add = function(line, js) {
        js? (code += line.match(reExp) ? line + '\n' : 'r.push(' + line + ');\n') :
            (code += line != '' ? 'r.push("' + line.replace(/"/g, '\\"') + '");\n' : '');
        return add;
    }
    while(match = re.exec(html)) {
        add(html.slice(cursor, match.index))(match[1], true);
        cursor = match.index + match[0].length;
    }
    add(html.substr(cursor, html.length - cursor));
    code += 'return r.join("");';
    return new Function(code.replace(/[\r\t\n]/g, '')).apply(options);
}

var template = 'My skills:' + 
    '<%for(var index in this.skills) {%>' + 
    '<a href="#"><%this.skills[index]%></a>' +
    '<%}%>';

console.log(TemplateEngine(template, {
    skills: ["js", "html", "css"]
}));

var template = 'My skills:' + 
'<%if(this.showSkills) {%>' +
    '<%for(var index in this.skills) {%>' + 
    '<a href="#"><%this.skills[index]%></a>' +
    '<%}%>' +
'<%} else {%>' +
    '<p>none</p>' +
'<%}%>';

console.log(TemplateEngine(template, {
    skills: ["js", "html", "css"],
    showSkills: true
}));
