// from Claude: doesn't work
function beep() {
  var context = new AudioContext();
  var oscillator = context.createOscillator();
  oscillator.type = 'sine';
  oscillator.frequency.setValueAtTime(1000, context.currentTime); 
  oscillator.start(0);
  
  setTimeout(function() {
    oscillator.stop(0);
  }, 200);  // 播放 200 毫秒
}



// from ChatGPT: it works
function beep() {
  // 创建一个新的AudioContext对象
  const AudioContext = window.AudioContext || window.webkitAudioContext;
  const audioContext = new AudioContext();

  // 创建一个OscillatorNode对象，它用于产生声音
  const oscillator = audioContext.createOscillator();

  // 将OscillatorNode连接到AudioContext的默认输出
  oscillator.connect(audioContext.destination);

  // 设置声音参数
  oscillator.type = 'square';  // sine/
  oscillator.frequency.value = 500;
  oscillator.start();


  // 创建一个GainNode对象，用于控制音量
  const gainNode = audioContext.createGain();
  oscillator.connect(gainNode);
  gainNode.connect(audioContext.destination);

  // 设置音量参数
  gainNode.gain.setValueAtTime(1, audioContext.currentTime);
  gainNode.gain.exponentialRampToValueAtTime(0.001, audioContext.currentTime + 1);

  // 在1000毫秒后停止声音
  setTimeout(() => {
    oscillator.stop();
  }, 100);
}

document.getElementById('root').onclick = function() { console.log("BEEP"); beep() }
