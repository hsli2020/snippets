// USPS = 420141329500110447760223484155
// USPS = 420141329202090153540046845346
// FEDEX = 9632013760767574185500395622321532
// CPC = 3LRL333010339000049296409000
// PURO = 1231831239460722803187000000002600

function cleanTracking(tn) {
    if (tn.length == 30) { // USPS
        return tn.substr(-22);
    }

    if (tn.length == 34) { // Fedex
        return tn.substr(-12);
    }

    if (tn.length == 28) { // CPC
        return tn.substr(7, 16);
    }

    if (tn.startsWith('14108')) { // PURO
        // tn ='1410810249460624659717700000002600'; // 606246597177
        tn = tn.substr(11, 12);
    } else if (tk.startsWith('12318')) {
        var n = tn.substr(15, 9);
        var a = parseInt(tn.substr(9,  2));
        var b = parseInt(tn.substr(11, 2));
        var c = parseInt(tn.substr(13, 2));
        tn = String.fromCharCode(64+a, 64+b, 64+c)+n;
    }

    return tn;
}

$('input[name=tracking]').blur(function() {
    var tn = $(this).val().trim();
    tn = cleanTracking(tn);
    $(this).val(tn);
});

function isPuro(tk) {
    return tk.startsWith('12318') || tk.startsWith('14108');
}

function isFedex(tk) {
    return tk.startsWith('9622') || tk.startsWith('9632');
}

function getPuroTK(tk) {
    if (tk.startsWith('14108')) {
      // tk ='1410810249460624659717700000002600'; // 606246597177
      return tk.substr(11, 12);
    } else if (tk.startsWith('12318')) {
      var n = tk.substr(15, 9);
      var a = parseInt(tk.substr(9,  2));
      var b = parseInt(tk.substr(11, 2));
      var c = parseInt(tk.substr(13, 2));
      return String.fromCharCode(64+a, 64+b, 64+c)+n;
    }
    return tk;
}

function getFedexTK(tk) {
    return tk.substr(-12);
}

$('.trim-tk').click(function() {
    var tk = $('input[name=tracking]');
    var tknum = tk.val();
    if (/\d{30,}/.test(tknum)) {
        if (isPuro(tknum)) {
            tk.val("PUR " + getPuroTK(tknum));
        } else if (isFedex(tknum)) {
            tk.val("Fedex " + getFedexTK(tknum));
        }
    }
})
