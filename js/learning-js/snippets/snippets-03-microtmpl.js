// Simple JavaScript Templating
// John Resig - http://ejohn.org/ - MIT Licensed
(function(){
  var cache = {};
 
  this.tmpl = function tmpl(str, data){
    // Figure out if we're getting a template, or if we need to
    // load the template - and be sure to cache the result.
    var fn = !/\W/.test(str) ?
      cache[str] = cache[str] ||
        tmpl(document.getElementById(str).innerHTML) :
     
      // Generate a reusable function that will serve as a template
      // generator (and which will be cached).
      new Function("obj",
        "var p=[],print=function(){p.push.apply(p,arguments);};" +

        // Introduce the data as local variables using with(){}
        "with(obj){p.push('" +
       
        // Convert the template into pure JavaScript
        str
          .replace(/[\r\t\n]/g, " ")
          .split("<%").join("\t")
          .replace(/((^|%>)[^\t]*)'/g, "$1\r")
          .replace(/\t=(.*?)%>/g, "',$1,'")
          .split("\t").join("');")
          .split("%>").join("p.push('")
          .split("\r").join("\\'")
      + "');}return p.join('');");
   
    // Provide some basic currying to the user
    return data ? fn( data ) : fn;
  };
})();
/*
You would use it against templates written like this 

    <script type="text/html" id="item_tmpl">
      <div id="<%=id%>" class="<%=(i % 2 == 1 ? " even" : "")%>">
        <div class="grid_1 alpha right">
          <img class="righted" src="<%=profile_image_url%>"/>
        </div>
        <div class="grid_6 omega contents">
          <p><b><a href="/<%=from_user%>"><%=from_user%></a>:</b> <%=text%></p>
        </div>
      </div>
    </script>

You can also inline script:

    <script type="text/html" id="user_tmpl">
      <% for ( var i = 0; i < users.length; i++ ) { %>
        <li><a href="<%=users[i].url%>"><%=users&#91;i&#93;.name%></a></li>
      <% } %>
    </script>

and you would use it from script like so:

    var results = document.getElementById("results");
    results.innerHTML = tmpl("item_tmpl", dataObject);
*/
