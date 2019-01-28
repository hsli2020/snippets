// func���û�������Ҫ�����ĺ���
// wait�ǵȴ�ʱ��
const debounce = (func, wait = 50) => {
  // ����һ����ʱ��id
  let timer = 0
  // ���ﷵ�صĺ�����ÿ���û�ʵ�ʵ��õķ�������
  // ����Ѿ��趨����ʱ���˾������һ�εĶ�ʱ��
  // ��ʼһ���µĶ�ʱ�����ӳ�ִ���û�����ķ���
  return function(...args) {
    if (timer) clearTimeout(timer)
    timer = setTimeout(() => {
      func.apply(this, args)
    }, wait)
  }
}
// ���ѿ�������û����øú����ļ��С��wait������£���һ�ε�ʱ�仹δ���ͱ�����ˣ�������ִ�к���

// �����������ȡ��ǰʱ�����
function now() {
  return +new Date()
}
/**
 * �������������غ�����������ʱ������ʱ�������ڻ���� wait��func �Ż�ִ��
 *
 * @param  {function} func        �ص�����
 * @param  {number}   wait        ��ʾʱ�䴰�ڵļ��
 * @param  {boolean}  immediate   ����Ϊtureʱ���Ƿ��������ú���
 * @return {function}             ���ؿͻ����ú���
 */
function debounce (func, wait = 50, immediate = true) {
  let timer, context, args

  // �ӳ�ִ�к���
  const later = () => setTimeout(() => {
    // �ӳٺ���ִ����ϣ���ջ���Ķ�ʱ�����
    timer = null
    // �ӳ�ִ�е�����£����������ӳٺ�����ִ��
    // ʹ�õ�֮ǰ����Ĳ�����������
    if (!immediate) {
      func.apply(context, args)
      context = args = null
    }
  }, wait)

  // ���ﷵ�صĺ�����ÿ��ʵ�ʵ��õĺ���
  return function(...params) {
    // ���û�д����ӳ�ִ�к�����later�����ʹ���һ��
    if (!timer) {
      timer = later()
      // ���������ִ�У����ú���
      // ���򻺴�����͵���������
      if (immediate) {
        func.apply(this, params)
      } else {
        context = this
        args = params
      }
    // ��������ӳ�ִ�к�����later�������õ�ʱ�����ԭ���Ĳ������趨һ��
    // �������ӳٺ��������¼�ʱ
    } else {
      clearTimeout(timer)
      timer = later()
    }
  }
}

������

/**
 * underscore �������������غ�����������ʱ��func ִ��Ƶ���޶�Ϊ �� / wait
 *
 * @param  {function}   func      �ص�����
 * @param  {number}     wait      ��ʾʱ�䴰�ڵļ��
 * @param  {object}     options   �������Կ�ʼ�����ĵĵ��ã�����{leading: false}��
 *                                �������Խ�β�����ĵ��ã�����{trailing: false}
 *                                ���߲��ܹ��棬����������ִ��
 * @return {function}             ���ؿͻ����ú���
 */
_.throttle = function(func, wait, options) {
    var context, args, result;
    var timeout = null;
    // ֮ǰ��ʱ���
    var previous = 0;
    // ��� options û������Ϊ�ն���
    if (!options) options = {};
    // ��ʱ���ص�����
    var later = function() {
      // ��������� leading���ͽ� previous ��Ϊ 0
      // �������溯���ĵ�һ�� if �ж�
      previous = options.leading === false ? 0 : _.now();
      // �ÿ�һ��Ϊ�˷�ֹ�ڴ�й©������Ϊ������Ķ�ʱ���ж�
      timeout = null;
      result = func.apply(context, args);
      if (!timeout) context = args = null;
    };
    return function() {
      // ��õ�ǰʱ���
      var now = _.now();
      // �״ν���ǰ�߿϶�Ϊ true
	  // �����Ҫ��һ�β�ִ�к���
	  // �ͽ��ϴ�ʱ�����Ϊ��ǰ��
      // �����ڽ��������� remaining ��ֵʱ�����0
      if (!previous && options.leading === false) previous = now;
      // ����ʣ��ʱ��
      var remaining = wait - (now - previous);
      context = this;
      args = arguments;
      // �����ǰ�����Ѿ������ϴε���ʱ�� + wait
      // �����û��ֶ�����ʱ��
 	  // ��������� trailing��ֻ������������
	  // ���û������ leading����ô��һ�λ�����������
	  // ����һ�㣬����ܻ���ÿ����˶�ʱ����ôӦ�ò��������� if ������
	  // ��ʵ���ǻ����ģ���Ϊ��ʱ������ʱ
	  // ������׼ȷ��ʱ�䣬�ܿ�����������2��
	  // ��������Ҫ2.2��Ŵ�������ʱ��ͻ�����������
      if (remaining <= 0 || remaining > wait) {
        // ������ڶ�ʱ����������������ö��λص�
        if (timeout) {
          clearTimeout(timeout);
          timeout = null;
        }
        previous = now;
        result = func.apply(context, args);
        if (!timeout) context = args = null;
      } else if (!timeout && options.trailing !== false) {
        // �ж��Ƿ������˶�ʱ���� trailing
	    // û�еĻ��Ϳ���һ����ʱ��
        // ���Ҳ��ܲ���ͬʱ���� leading �� trailing
        timeout = setTimeout(later, remaining);
      }
      return result;
    };
  };