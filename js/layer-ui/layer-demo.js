/*! layer demo */

;!function(){

var gather = {
  htdy: $('html, body')
};

//全局配置
layer.config({
  extend: 'skin/moon/style.css'
});


//一睹为快
gather.demo1 = $('#demo1');
$('#chutiyan>a').on('click', function(){
  var othis = $(this), index = othis.index();
  var p = gather.demo1.children('p').eq(index);
  var top = p.position().top;
  gather.demo1.animate({scrollTop: gather.demo1.scrollTop() + top}, 0);
  switch(index){
    case 0:
      var icon = -1;
      (function changeIcon(){
        var index = layer.alert('Hi，你好！ 点击确认更换图标', {
          icon: icon,
          shadeClose: true,
          title: icon === -1 ? '初体验 - layer '+layer.v : 'icon：'+icon + ' - layer '+layer.v
        }, changeIcon);
        if(8 === ++icon) layer.close(index);
      }());
    break;
    case 1:
      var icon = 0;
      (function changeIcon1(){
        var index = layer.alert('点击确认更换图标', {
          icon: icon,
          shadeClose: true,
          skin: 'layer-ext-moon',
          shift: 5,
          title: icon === -1 ? '第三方扩展皮肤' : 'icon：'+icon
        }, changeIcon1);
        if(9 === ++icon) {
          layer.confirm('怎么样，是否很喜欢该皮肤，去下载？', {
            skin: 'layer-ext-moon'
          }, function(index, layero){
            layero.find('.layui-layer-btn0').attr({
              href: 'http://layer.layui.com/skin.html',
              target: '_blank'
            });
            layer.close(index);
          });
        };
      }());
    break;
    
    /*
    case 5:
      layer.open({
        type: 1,
        shade: false,
        title: false, //不显示标题
        content: $('.layer_notice'), //捕获的元素
        cancel: function(index){
          layer.close(index);
          this.content.show();
          layer.msg('捕获就是从页面已经存在的元素上，包裹layer的结构',{icon:6, time: 5000});
        }
      });
      $('#sougouAD').addClass('popupad');
    break;
    */
    
    case 6:
      layer.open({
        type: 1,
        area: ['420px', '240px'],
        skin: 'layui-layer-rim', //加上边框
        content: '<div style="padding:20px;">即直接给content传入html字符<br>当内容宽高超过定义宽高，会自动出现滚动条。<br><br><br><br><br><br><br><br><br><br><br>很高兴在下面遇见你</div>'
      });
    break;
    case 7:
      layer.open({
        type: 1,
        skin: 'layui-layer-demo',
        closeBtn: false,
        area: '350px',
        shift: 2,
        shadeClose: true,
        content: '<div style="padding:20px;">即传入skin:"样式名"，然后你就可以为所欲为了。<br>你怎么样给她整容都行<br><br><br>我是华丽的酱油==。</div>'
      });
    break;
    case 8:
      layer.tips('Hi，我是tips', this);
    break;
    case 11:
      var ii = layer.load(0, {shade: false});
      setTimeout(function(){
        layer.close(ii)
      }, 5000);
    break;
    case 12:
      var iii = layer.load(1, {
        shade: [0.1,'#fff']
      });
      setTimeout(function(){
        layer.close(iii)
      }, 3000);
    break;
    case 13:
      layer.tips('我是另外一个tips，只不过我长得跟之前那位稍有些不一样。', this, {
        tips: [1, '#3595CC'],
        time: 4000
      });
    break;
    case 14:
      layer.prompt({title: '输入任何口令，并确认', formType: 1}, function(pass){
        layer.prompt({title: '随便写点啥，并确认', formType: 2}, function(text){
          layer.msg('演示完毕！您的口令：'+ pass +'<br>您最后写下了：'+text);
        });
      });
    break;
    case 15:
      layer.tab({
        area: ['600px', '300px'],
        tab: [{
          title: '无题', 
          content: '<div style="padding:20px; line-height:30px; text-align:center">欢迎体验layer.tab<br>此时此刻不禁让人吟诗一首：<br>一入前端深似海<br>从此妹纸是浮云<br>以下省略七个字<br>。。。。。。。<br>——贤心</div>'
        }, {
          title: 'TAB2', 
          content: '<div style="padding:20px;">TAB2该说些啥</div>'
        }, {
          title: 'TAB3', 
          content: '<div style="padding:20px;">有一种坚持叫：layer</div>'
        }]
      });
    break;
    case 16:
      if(gather.photoJSON){
        layer.photos({
          photos: gather.photoJSON
        });
      } else {
        $.getJSON('test/photos.json?v='+new Date, {}, function(json){
          gather.photoJSON = json;
          layer.photos({
            photos: json
          });
        });
      }
    break;
    default:
      new Function(p.text())();
    break;
  }
});

//一往而深
$('#demore').on('click', function(){
  gather.htdy.animate({scrollTop : $('#yiwang').offset().top}, 200);
});
gather.demo2 = $('#demo2');
$('.layer-demolist').on('click', function(){
  var othis = $(this), index = othis.index('.layer-demolist');
  var p = gather.demo2.children('p').eq(index);
  var top = p.position().top;
  gather.demo2.animate({scrollTop: gather.demo2.scrollTop() + top}, 0);
  switch(index){
    case 15:
      layer.tips('上', this, {
        tips: [1, '#000']
      });
    break;
    case 16:
      layer.tips('默认就是向右的', this);
    break;
    case 17:
      layer.tips('下', this, {
        tips: 3
      });
    break;
    case 18:
      layer.tips('在很久很久以前，在很久很久以前，在很久很久以前……', this, {
        tips: [4, '#78BA32']
      });
    break;
    case 19:
      layer.tips('不会销毁之前的', this, {tipsMore: true});
    break;
     default:
      new Function(p.text())();
    break;
  }
});

//异步请求
gather.downs = $('#downs');
gather.downs [0] && function(){
  
  //获取下载数
  $.get('http://fly.layui.com/api/handle?id=1&type=find', function(res){
    gather.downs.html(res.number);
  }, 'jsonp');

  
  //获取并记录关注次数
  $.get('http://fly.layui.com/api/handle?id=3', function(res){
    $('#sees').html(res.number);
  }, 'jsonp');
}();

//记录下载
$('.layer_down').on('click',function(){
  $.get('http://fly.layui.com/api/handle?id=1');
});

//API页
gather.api = $('.layer-api');
gather.apiRun = $('.layer-api-run');
(function(){
  var lis = gather.api.find('li'), slecked = 'layer-api-slecked';
  lis.on('click', function(){
    lis.removeClass(slecked);
    $(this).addClass(slecked);
  });
  gather.api.find('h2').on('click', function(){
    var othis = $(this), i = othis.find('.layer-api-ico');
    if(i.hasClass('icon-shousuo')){
      i.addClass('icon-zhankai').removeClass('icon-shousuo');
      othis.next().hide();
      
    } else {
      i.addClass('icon-shousuo').removeClass('icon-zhankai');
      othis.next().show();
    }
  });
  layer.ready(function(){
    layer.photos({
      photos: '#layer-photos-demo'
    });
  });
}());

gather.skin = function(){
  var index = layer.open({
    type: 1,
    title: 'layer皮肤制作说明',
    skin: 'layer-ext-moon',
    area: '888px',
    content: $('#skinFabu'),
    shadeClose: true,
    btn: ['不大明白', '我知道了'],
    yes: function(){
      layer.confirm('OMG，是我的文档写的太烂了吗？好吧，是否去看看别人是如何制作的？', {
        icon: 8
      }, function(){
        location.href = 'http://layer.seaning.com';
      }, function(){
        layer.closeAll();
      });
    }
  });
  layer.full(index);
}

//发布皮肤
gather.pub = $('#skinPublish');
gather.pub.on('click', gather.skin);
if(gather.pub[0] && location.hash === '#publish'){
  layer.ready(function(){
    gather.skin();
  });
}

//窗口scroll
(function(){
  var conf = {};
  //返回顶部
  conf.gotop = $('.layer_gotop');
  conf.gotop.on('click',function(){
    $('html, body').animate({scrollTop : 0},$(this).offset().top/7);
  });
  conf.log = 0;
  $(window).on('scroll', function(){
    var stop = $(window).scrollTop();
    if(stop >= 60){
      if(!conf.log){
        conf.log = 1;
        conf.gotop.show();
        gather.api.addClass('layer-api-fix');
        gather.apiRun.css('top', 0);
      }
    } else {
      if(conf.log){
        conf.log = 0;
        conf.gotop.hide();
        gather.api.removeClass('layer-api-fix');
        gather.apiRun.css('top', '60px');
      }
    }
    stop = null;
  });
}());

//修饰代码
$('pre').each(function(i){
  var othis = $(this);
  othis.show().laycode({
    title: othis.attr('title') || '对应代码说明',
    height: othis.attr('heg') || 'auto',
    skin: othis.attr('skin') || 0
  });
});


//ie6
if(!-[1,] && !window.XMLHttpRequest){
  layer.ready(function(){
    layer.alert('如果您是用ietest的ie6模式，发现弹出背景一片黑色时，请不用惊慌，这并非layer未作兼容，而是你当前版本的ietest所模拟的ie6环境未对滤镜做支持，标准ie6将不会有此问题，所以请您不要担心。');
  });
}

gather.getDate = function(time){ 
  return new Date(parseInt(time)).toLocaleString() 
};

//众筹
gather.chou = function(){
  layer.alert('感谢过去一个月里热心参与活动的朋友们，目前六本《JavaScript权威指南》已全部送出。<br><br>请不要再通过任何渠道给予赞助，layui自始自终坚持无偿服务！<p style="text-align:right;">2015.07.07</p>')
  /*
  var index = layer.open({
    type: 1
    ,title: '赞助额前六名的朋友，将回赠《Javascript权威指南·第六版》（￥99.1）'
    ,area: ['100%', '100%']
    ,scrollbar: false
    ,move: false
    ,shift: 2
    ,skin: 'layer-ext-moon'
    ,end: function(){
      location.hash = '';
    }
    ,content: '<div class="layer-chou">\
      <p>自始自终，都并不希望善良的您为layer破费一分钱，甚至包括后续发布的layui，永远都会保持无偿服务的初心、为web开发提供强劲动力的使命不会改变。</p>\
      <p>我总是在期盼着我虔诚编织的这一切，都能够融入到越来越多人的工作中，能够真正地为我们这个行业，为这个社会，这个国家贡献一些尽可能多的能量。</p>\
      <p>然而，一个人的力量始终有限，在我们还并不富有的时候，这些梦想的支撑都需要一定精神与物质的基础。因此我终于厚着脸皮，艰难地开启了这个众筹。</p>\
      <p style="margin-top:10px;">如果您感受到了这份信任，那么可用支付宝扫描下面的二维码，给予赞助。</p>\
      <img style="width:150px; height:150px;" src="http://res.layui.com/images/pay.png">\
      <p>在活动结束的<em style="color:#999">7月7日</em>，金额前<em>六</em>名的朋友，将会获得此书：<a href="http://product.dangdang.com/22722790.html" target="_balnk" style="color:#4285F4">《Javascript权威指南·第六版》</a>，每本<em style="color:#FF4400">99元</em>。</p>\
      <p>而剩下的资金，将全部用于layui的服务器费用</p>\
      <p style="margin-top:10px;">如果您已经赞助，为了和支付结果匹配，请务必填写以下信息。</p>\
      <ul id="C_ul">\
        <li><label>姓名或网名：</label><input id="C_name"></li>\
        <li><label>已赞助金额：</label><input id="C_sum" type="number"></li>\
        <li><label>您的QQ号：</label><input id="C_qq" type="number"> <span>*便于与您取得联系，不会公开</span></li>\
        <li><label> </label><a id="C_btn" class="btns" href="javascript:;">提交</a></li>\
      </ul>\
      <dl id="C_list"></dl>\
    </div>'
  });
  //layer.full(index);
  
  location.hash = 'donate';
  
  $('#C_btn').on('click', function(){
    var data = {
      name: $('#C_name').val()
      ,sum: $('#C_sum').val()
      ,qq: $('#C_qq').val()
    };
    if(data.name.length < 2 || data.sum === '' || data.qq.length < 5){
      layer.msg('严肃点');
      return;
    }
    $.ajax({
       url: 'http://layms.layui.com/ajax/chou',
       dataType: 'jsonp',
       data: data,
       success: function(res){
         if(!res.status){
           layer.msg('提交成功，匹配Ok后会在下放给予显示', {icon: 1});
           $('#C_ul').remove();
           return;
         };
         layer.msg('提交失败');
       }, error: function(){
         layer.msg('请求异常');
       }
    });
  });
  
  //获取捐赠名单
  $.ajax({
    url: 'http://layms.layui.com/ajax/chou?type=1',
    dataType: 'jsonp',
    success: function(res){
      if(!res.status){
        var data = res.data, str = '<dt>赞助额TOP10</dt>', len = data.length;
        if(len === 0){
          str = '<dd style="color:#999">暂无赞助记录</dd>';
        }
        for(var i = 0; i < data.length; i++){
          str += '<dd>'+ data[i].name + '：<em>'+ data[i].sum +'￥</em>（'+ gather.getDate(data[i].time) +'）</dd>'
        }
        $('#C_list').html(str);
      }
    }
  });
  */
  return index;
}
$('#chou').on('click', gather.chou);
if(location.hash === '#donate'){
  layer.ready(function(){
    gather.chou();
  });
}

//弹出广告
$('.ads').on('click', function(){
   $(this).removeClass('popupad'); 
});


window.paysentsin = function(){
  return layer.photos({
    photos: {
      "title": ""
      ,"data": [
        {
          "alt": "layer友情打赏",
          "pid": 666, //图片id
          "src": "http://res.layui.com/images/pay/layer.jpg", //原图地址
          "thumb": "" //缩略图地址
        }
      ]
    }
  });
};

//判断是iframe打开页面
if(top.location.host !== 'layer.layui.com' && top.location != self.location){
  layer.ready(function(){
    layer.alert('由于百度抽风，layer官网突然降权，您当前进入的不是官网。<br>请浏览器收藏唯一域名：<a href="http://layer.layui.com" target="_blank" style="color:#5fbfe7">http://layer.layui.com</a>', {
      offset: '1px', 
      btn: '去官网',
      closeBtn: false,
      move: false,
      maxWidth: 600,
      success: function(layero){
        layero.find('.layui-layer-btn a').attr({
          href: 'http://layer.layui.com',
          target: '_blank'
        });
      }
    }, function(){
      
    });
  });
}
  
}();