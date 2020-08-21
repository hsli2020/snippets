// https://github.com/ppz-pro/noty.js

// index.js

import Noty from './noty.js'
import CreateAnimation from './animation/index.js'
import CreateContent from './content/index.js'

export default Noty({
  Animation: CreateAnimation(),
  Content: CreateContent()
})

// noty.js

export default function Noty({
  rootClass = '',
  customPosition,
  mount = document.querySelector('body'),
  Content,
  Animation,
  duration = 3000
}) {
  const root = document.createElement('div')
  root.className = 'ppz-noty-root ' + rootClass
  if(!customPosition) {
    root.style.position = 'fixed'
    root.style.right = '1em'
    root.style.top = '1em'
  }
  mount.append(root)

  return new Proxy(Content, {
    get(target, method) {
      return function() {
        let closed = false
        const instance = {
          root,
          duration, // 传到“目标方法”里，便于设置 duration（show 之前设置有效，之后无效）
          closed: () => closed, // 防止外部修改
          close() {
            if(closed) throw Error('通知已经关闭')
            if(instance.onClose) instance.onClose()
            closed = true
            Animation.close(content, root)
          }
        }
        var content = instance.content = Content[method].apply(instance, arguments)
        Animation.show(content, root)
        if(instance.duration != 0)
          instance.timeoutID = setTimeout(instance.close, instance.duration)
        return instance
      }
    }
  })
}

// content/index.js 

// 想用这个 content 的样式，只引用 css 就可以了

export default function CreateContent() {
  return {
    info: createShow('info'),
    success: createShow('success'),
    warn: createShow('warn'),
    error: createShow('error'),
    loading: createShow('loading', { duration: 0 })
  }
}

function createShow(className, defaults) {
  return function(text, options = {}) {
    if(defaults)
      Object.assign(options, defaults)
    if(options.duration != undefined)
      this.duration = options.duration
    this.onClose = options.onClose
    
    const div = Div('', // content 部分再包一层，避免之上的样式影响动画效果
      `<div class='ppz-noty-item1 ppz-noty-item1-${className}'>
        <i class='ppz-noty-icon ppz-noty-icon-${className}'></i>
        <span class='text'></span>
      </div>`
    )
    div.querySelector('.text').append(text) // 使用 append 方法，省去 escape 操作
    if(options.closeBtn) {
      const btn = Div('close-btn', 'x')
      btn.onclick = this.close
      div.querySelector('.ppz-noty-item1').prepend(btn)
    }
    return div
  }
}

function Div(className, innerHTML) {
  const div = document.createElement('div')
  div.className = className
  div.innerHTML = innerHTML
  return div
}

// animation/index.js

export default function CreateAnimation({ duration = 300} = {}) { // 动画持续时间
  const list = [] // 存放所有未关闭的 content（用于区别“关闭后”但动画未结束而滞留在文档流里的 content）
  return {
    show(content, container) {
      list.push(content) // 加入 list
      container.append(content) // 加入文档流
      content.animate([ // 动画
        {
          transform: 'translateY(100%)',
          opacity: 0
        },
        {}
      ], duration)
    },
    close(content) {
      // FILP 原理 https://aerotwist.com/blog/flip-your-animations/
      // first: 原来的位置
      const firsts = list.map(content => content.getBoundingClientRect())
      // last:
      content.style.position = 'fixed' // (1)从文档流中移除目标 content
      const lasts = list.map(content => {
        content.getAnimations().forEach(ani => ani.cancel()) // (2)取消（停止）一切动画（得到最终“形态”）（同时触发下面的 oncancel）
        return content.getBoundingClientRect() // 1 + 2 => 得到 last
      })
      
      list.forEach((ctt, index) => {
        const first = firsts[index]
        ctt.style.left = first.left + 'px'
        if(ctt == content) { // 待关闭的 content，只要往上移动一个身位就行了（不需要 last）
          ctt.animate([
            { top: first.top + 'px' }, // 回到 first（此时处于“最终形态”）
            {
              top: first.top - first.height + 'px', // 往上移一个身位
              opacity: 0 // fadeOut
            }
          ], duration)
          .onfinish = function() {
            ctt.remove() // 动画结束时，移出文档流
          }
        } else {
          const last = lasts[index]
          ctt.style.position = 'fixed'
          const ani = ctt.animate([
            { top: first.top + 'px' }, // 回到：first
            { top: last.top + 'px' } // 去向：last
          ], duration)
          ani.onfinish = ani.oncancel = function() {
            ctt.style.position = '' // 立刻回归**真实**位置
          }
        }
      })
      // 立刻从 list 中移出 content（此时 content 还因动画未结束而滞留在文档流）
      list.splice(list.indexOf(content), 1)
    }
  }
}
