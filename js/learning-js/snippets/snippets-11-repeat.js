String.prototype.repeat= function(n){
    return Array(n+1).join(this);
}

function repeat(string, times) {
  var result = "";
  for (var i = 0; i < times; i++)
    result += string;
  return result;
}

