//---------------------------------------------------------
// 简单的Toast (可以改进一下让它能同时显示多个)

<div class="alert-info" style="display:none;"><span></span></div>

<script type="text/javascript">
var Toast = (function() {
    "use strict";

    var elem, that = {};

    that.init = function(options) {
        elem = $(options.selector);
    };

    that.show = function(text) {
        elem.find("span").html(text);
        elem.delay(200).fadeIn().delay(4000).fadeOut();
    };

    return that;
}());

$(document).ready(function() {
    Toast.init({selector: '.alert-info'});
    //Toast.show('Toast is ready for you');
}());
</script>
        
<style>
.alert-info {
    position: fixed;
    top: 5%;
    right: 10px;
    z-index: 2000;
    background-color: #3b686d;
    border-color: #bce8f1;
    color: white;
    border: 1px solid transparent;
    border-radius: 4px;
    padding: 10px 15px;
}
</style>

//---------------------------------------------------------
// 另外一种风格

  <script type="text/javascript">
    $(document).ready(function(){
        $('button').click(function () {
            var msg = $(this).data('msg');
            $('.error').text(msg).fadeIn(400).delay(3000).fadeOut(400); 
        });
    }); 
  </script>
  
  <style type="text/css">
    .error {
      /*width:200px;*/
        height:20px;
        height:auto;
        position:absolute;
        left:50%;
        margin-left:-100px;
        bottom:10px;
        background-color: #383838;
        color: #F0F0F0;
        font-family: Calibri;
        font-size: 20px;
        padding:10px;
        text-align:center;
        border-radius: 2px;
        -webkit-box-shadow: 0px 0px 24px -1px rgba(56, 56, 56, 1);
        -moz-box-shadow: 0px 0px 24px -1px rgba(56, 56, 56, 1);
        box-shadow: 0px 0px 24px -1px rgba(56, 56, 56, 1);
    }
  </style>

<button data-msg='I did something good better best!'>Do something!</button>

<div class='error' style='display:none'></div>
