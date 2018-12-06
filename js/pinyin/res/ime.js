// Gecko = Netscape 7.x & Mozilla 1.4+

LastNo=0;
SPACECHAR=" ";
CandChinesePart=new Array();
CandCompPart=new Array();
AsciiStr="a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z";
AsciiStr=AsciiStr.split(',');
CodeList=CodeList.split(',');

CtrlDown=false;
CancelKey=false;

//==KeyCode 33..47 
//Symbol1 = "！＂＃＄％＆＇（）＊＋，－。／"; 
//==KeyCode 58..64 
//Symbol2 = "：；＜＝＞？＠"; 
//==KeyCode 91..96 
//Symbol3 = "［、｝＾＿｀"; 
//==KeyCode 123..126 
//Symbol4 = "｛｜］～"; 

Punct2 = new Array('；','＝','，','－','。','／','｀'); 
Punct3 = new Array('［','、','］','＇'); 

SPunct1 = new Array('）','！','＠','＃','＄','％','＾','＆','＊','（'); 
SPunct2 = new Array('：','＋','＜','＿','＞','？','～'); 
SPunct3 = new Array('｛','｜','｝','＂'); 

FullShape_No=new Array("０","１","２","３","４","５","６","７","８","９"); 
FullShape_BigAZ=new Array("Ａ","Ｂ","Ｃ","Ｄ","Ｅ","Ｆ","Ｇ","Ｈ","Ｉ","Ｊ","Ｋ","Ｌ","Ｍ","Ｎ","Ｏ","Ｐ","Ｑ","Ｒ","Ｓ","Ｔ", "Ｕ","Ｖ","Ｗ","Ｘ","Ｙ","Ｚ"); 
FullShape_SmallAZ=new Array('ａ','ｂ','ｃ','ｄ','ｅ','ｆ','ｇ','ｈ','ｉ','ｊ','ｋ','ｌ','ｍ','ｎ','ｏ','ｐ','ｑ','ｒ','ｓ','ｔ','ｕ','ｖ','ｗ','ｘ','ｙ','ｚ'); 

function FindIn(s) {
var find=-1,low=0,mid=0,high=CodeList.length;
var sEng="";
  while(low<high){
    mid=(low+high)/2;
    mid=Math.floor(mid);
    sEng=CodeList[mid];
    if(sEng.indexOf(s,0)==0){find=mid;break;}
    sEng=CodeList[mid-1];
    if(sEng.indexOf(s,0)==0){find=mid;break;}
    if(sEng<s){low=mid+1;}else{high=mid-1;}
  }
  while(find>0){
    sEng=CodeList[find-1];
    if(sEng.indexOf(s,0)==0){find--;}else{break;}
  }
  return(find);
}

function GetStr(No, s){
  var sTmp="",sChi="";
  var i;
  for(i=0;i<=9;i++){
    if(No+i>CodeList.length-1){break;}
    sTmp=CodeList[No+i];
    if(sTmp.indexOf(s)==0){
      sChi=CodeList[No+i];
      CandCompPart[i]=sChi.substring(s.length,sChi.indexOf(SPACECHAR));
      CandChinesePart[i]=sChi.substr(sChi.lastIndexOf(SPACECHAR)+1);
      if(i<=8){IME.Cand.value+=(i+1)+"."+CandChinesePart[i]+CandCompPart[i]+'\n';}
      else{IME.Cand.value+=(0)+"."+CandChinesePart[i]+CandCompPart[i]+'\n';}
    }else{
      break;
    }
  }
  if(No>10 && CodeList[No-10].indexOf(s)==0) IME.Cand.value+='-.<<\n';
  if(i==10 && No<=CodeList.length-11 && CodeList[No+i].indexOf(s)==0) IME.Cand.value+='+.>>';
  LastNo=No;
}

function Grep(s){
  var No=-1;
  for(i=0;i<=9;i++){CandChinesePart[i]="";}
  if(s!=""){
    No=FindIn(s);
    if(No>=0){GetStr(No, s);}
  }

  if( IME.AutoUp.checked==true && CandChinesePart[0]!="" && CandChinesePart[1]=="" && CandCompPart[0]=="" ) {
    SendCand(0);
  }
}

function SendCand(n){
  if ( n>=0 && n<=9 ) {
    SendStr(CandChinesePart[n]);
    IME.Comp.value="";
    IME.Cand.value="";
  }
}

function setSelectionRange(input, selectionStart, selectionEnd) {
  if (input.setSelectionRange) {
    input.focus();
    input.setSelectionRange(selectionStart, selectionEnd);
  } else if (input.createTextRange) {
    var range = input.createTextRange();
    range.collapse(true);
    range.moveEnd('character', selectionEnd);
    range.moveStart('character', selectionStart);
    range.select();
  }
}

function setCaretToEnd (input) {
  setSelectionRange(input, input.value.length, input.value.length);
}
function setCaretToBegin (input) {
  setSelectionRange(input, 0, 0);
}
function setCaretToPos (input, pos) {
  setSelectionRange(input, pos, pos);
}


function SendStr(s) {
  if (s=="") { return }

  switch (browser) {
  case 1: // IE
    var r=document.selection.createRange();
    r.text=s;
    r.select();
    break;
  case 2: // Gecko
    /*
       -simulate keypress
       *simulate scroll
       -change outputString
       -use createRange, setStart, setEnd
    */

    var obj = IME.InputArea;
    var selectionStart = obj.selectionStart;
    var selectionEnd = obj.selectionEnd;
    var oldScrollTop = obj.scrollTop;
    var oldScrollHeight = obj.scrollHeight;
    var oldLen = obj.value.length;

    obj.value = obj.value.substring(0, selectionStart) + s + obj.value.substring(selectionEnd);
    obj.selectionStart = obj.selectionEnd = selectionStart + s.length;
    if (obj.value.length == oldLen) {
      obj.scrollTop = obj.scrollHeight;
    } else if (obj.scrollHeight > oldScrollHeight) {
      obj.scrollTop = oldScrollTop + obj.scrollHeight - oldScrollHeight;
    } else {
      obj.scrollTop = oldScrollTop;
    }
    
    break;
  default:
    IME.InputArea.value += s;
  }
}


function ToFullShapeLetter(aStr) {
  var s="";

  for (i=0;i<aStr.length;i++) {
    var c = aStr.charCodeAt(i);
    if (c>=65 && c<=90) {
      s += FullShape_BigAZ[c-65];
    } else if (c>=97 && c<=122) {
      s += FullShape_SmallAZ[c-97];
    } else {
      s += aStr.charAt(i);
    }
  }

  return s;
}

function ImeKeyDown(e) {
  var s="";
  if(!e) e=window.event;
  var key = e.which ? e.which : e.keyCode;
  CtrlDown=false;
  passNextKeyPress=false;
  
  switch (key) {
  //==Backspace
  case 8:
    if (IME.Comp.value!="") {
      s=IME.Comp.value;
      IME.Comp.value = s.substr(0, s.length-1);
      IME.Cand.value = "";
      Grep(IME.Comp.value);
      CancelKey = true;
      return (false);
    }
    return (true);

  //==Tab
  case 9:
    SendStr('¡¡');
    CancelKey = true;
    return (false); 

  //==Esc
  case 27:
    IME.Comp.value="";
    IME.Cand.value="";
    CancelKey = true;
    return (false);

  //
  case 109: //firfox keycode '-'
       if(browser != 2) break;
  case 189: //ie keycode '-'
       if(browser != 1 && key != 109) break;
  //==PageUp
  case 33:
  case 57383: // Opera 7.11
    s=IME.Comp.value;
    if (s!="") {
      if(LastNo>10 && CodeList[LastNo-10].indexOf(s)==0){
        IME.Cand.value="";
        GetStr(LastNo-10, s);
      }
      CancelKey = true;
      return(false);
    }
    break;
    //return(true);

  case 61: //firfox keycode '-'
       if(browser != 2) break;
  case 187: //ie keycode '-'
       if(browser != 1 && key != 61) break;
  //==PageDown
  case 34:
  case 57384: // Opera 7.11
    s=IME.Comp.value;
    if ( s!="" ){
      if(LastNo<=CodeList.length-11 && CodeList[LastNo+10].indexOf(s) == 0) {
        IME.Cand.value="";
        GetStr(LastNo+10, s);
      }
      CancelKey = true;
      return(false);
    }
    break;
    //return(true);

  //==Space
  case 32:
    if(IME.Comp.value!="") {
      //TODO: sound if nothing in Cand
      SendCand(0);
      CancelKey = true;
      return(false);
    }
    return(true);

  //==Enter
  case 13:
    if (IME.Comp.value!="") {
      SendStr( IME.FullShape.checked ? 
        ToFullShapeLetter(IME.Comp.value) : 
        IME.Comp.value);
      IME.Comp.value="";
      IME.Cand.value="";
      CancelKey = true;
      return(false);
    }
    return(true);

  //==F2
  case 113:
    IME.AutoUp.checked = !IME.AutoUp.checked;
    CancelKey = true;
    return (false);

  //==F12
  case 123:
  case 57356: // Opera 7.11
    IME.FullShape.checked = !IME.FullShape.checked;
    CancelKey = true;
    return (false);

  //==Ctrl
  case 17:
  case 57402: // Opera 7.11
    CtrlDown=true;
    break;
    
  		case 36: // home
		case 35: // end
		case 37: // left
		case 38: // up
		case 39: // right
		case 40: // down
		case 45: // insert
		case 46: // del
		case 91: // windows key
		case 112: // F1
//		case 113: // F2
		case 114: // F3
		case 115: // F4
		case 116: // F5
		case 117: // F6
		case 118: // F7
		case 119: // F8
		case 120: // F9
		case 121: // F10
		case 122: // F11
//		case 123: // F12
			// let these keys pass through unprocessed in the next keypress events
			passNextKeyPress = true;
			return (true);

  }

  if (e.ctrlKey) { return (true) };


  if (key>=48 && key<=57) {
    if (e.shiftKey) {
      if (IME.FullShape.checked || !IME.EnglishMode.checked) {
        SendStr(SPunct1[key-48]);
        CancelKey = true;
        return (false);
      }
    } else {
      if (IME.Comp.value=="") {
        if (IME.FullShape.checked || !IME.EnglishMode.checked) {
          SendStr(FullShape_No[key-48]);
          CancelKey = true;
          return (false);
        }
      } else {
        if (!IME.EnglishMode.checked) {
          SendCand( key==48 ? 9 : (key-49) );
          CancelKey = true;
          return (false);
        }
      }
    }
    return (true);
  }

  if (IME.FullShape.checked || !IME.EnglishMode.checked) {
    if (key>=186 && key<=192) {
      SendStr( e.shiftKey ? SPunct2[key-186] : Punct2[key-186] );
      CancelKey = true;
      return (false);
    }
    if (key>=219 && key<=222) {
      SendStr( e.shiftKey ? SPunct3[key-219] : Punct3[key-219] );
      CancelKey = true;
      return (false);
    }
  }    

  if (browser==2) {
    if (IME.FullShape.checked || !IME.EnglishMode.checked) {
      switch (key) {
      case 59:
        SendStr( e.shiftKey ? SPunct2[0] : Punct2[0] );
        CancelKey = true;
        return (false);
      case 61:
        SendStr( e.shiftKey ? SPunct2[1] : Punct2[1] );
        CancelKey = true;
        return (false);
      case 109:
        SendStr( e.shiftKey ? SPunct2[3] : Punct2[3] );
        CancelKey = true;
        return (false);
      }
    }
  }

  return(true);
}

function ImeKeyPress(e) {
  if(!e) e=window.event;
  var key = e.which ? e.which : e.keyCode;

	// pass keypress without processing it
	if(passNextKeyPress) {
		return (true);
	}

  if (browser==2 || browser==3) {
    if (CancelKey) {
      CancelKey = false;
      return (false);
    }
  }

  if (e.ctrlKey) { return (true); }

  //==A-Z
  if ( key>=65 && key<=90 ) {
    if (IME.FullShape.checked) {
      SendStr(FullShape_BigAZ[key-65]);
      return (false);
    }
    return (true);
  }

  //==a-z
  if (key>=97 && key<=122) {
    if (IME.EnglishMode.checked) {
      if (IME.FullShape.checked) {
        SendStr( FullShape_SmallAZ[key-97] );
        return (false);
      }
      return (true);
    } else {
      s=IME.Comp.value;
      if (s.length<MAX) {
        IME.Comp.value+=AsciiStr[key-97];
        IME.Cand.value="";
        Grep(IME.Comp.value);
      }
      return (false);
    }
  }

  if (browser==2) {
    switch (key) {
    case 47: case 63:
      if (!IME.EnglishMode.checked || IME.FullShape.checked) {
        return (false);
      }
      break;
    }
  }

  return (true);
}

function ImeKeyUp(e) {
  if(!e) e=window.event;
  var key = e.which ? e.which : e.keyCode;

  //==Ctrl
  if (key==17 || key==57402) {
    if (CtrlDown==true) {
      IME.EnglishMode.checked = !IME.EnglishMode.checked;
    }
  }

  return(true);
}

function BodyOnLoad() {
  browser = 
    (navigator.appName.indexOf('Microsoft') != -1) ? 1 :
    (navigator.appName.indexOf('Netscape')  != -1) ? 2 :
    (navigator.appName.indexOf('Opera')     != -1) ? 3 :
    4;
  if(browser == 2 && navigator.userAgent.indexOf('Safari') != -1) browser =
5;

  IME = {
    InputArea: document.getElementById("InputArea"),
    Comp:      document.getElementById("Comp"),
    Cand:      document.getElementById("Cand"),

    EnglishMode: document.getElementById("EnglishMode"),
    FullShape:   document.getElementById("FullShape"),
    AutoUp:      document.getElementById("AutoUp")
  }
  IME.InputArea.focus();
  
}

//BETA
function LoadImeTable() {}


