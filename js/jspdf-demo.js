https://pastebin.com/gPZDXRqF


var script = document.createElement('script');
script.type = 'text/javascript';
script.src = 'https://cdnjs.cloudflare.com/ajax/libs/jspdf/2.2.0/jspdf.umd.min.js';
document.head.appendChild(script);


let cs = document.querySelectorAll('.pages-content canvas');
let doc = new jspdf.jsPDF('p', 'pt', 'a4', true);
let w = doc.internal.pageSize.width, h = doc.internal.pageSize.height;
for (var i=0; i < cs.length; i++) { if (i > 0) doc.addPage(); await doc.addImage(cs[i], 'png', 0, 0, w, h); }
doc.save("output.pdf")

