<script>
  (function () {
    var s = document.createElement('script');
    s.type = 'text/javascript';
    s.async = true;
    s.src = 'http://xxx.com/sdk.js';
    var x = document.getElementsByTagName('script')[0];
    x.parentNode.insertBefore(s, x);
  })();
</script>

针对现代浏览器，可以使用async。
<script async src="http://xxx.com/sdk.js"></script>

传统语法
<script type="text/javascript" src="http://xxx.com/sdk.js"></script>

