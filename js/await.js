//console.log('before fetch');

/*
fetch('http://play.dev/index.php', {
	method: 'get'
}).then(function(response) {
    return response.text();
}).then(function(text) {
    console.log(text);
}).catch(function(err) {
	console.log('Error');
});
*/

async function play() {
    console.log('before fetch');
    var res = await fetch('http://play.dev/index.php');
    var txt = await res.text();
    console.log(txt);
    console.log('after fetch');
}
play();
