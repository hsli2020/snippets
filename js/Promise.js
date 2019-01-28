// ����״̬
const PENDING = "pending";
const RESOLVED = "resolved";
const REJECTED = "rejected";
// promise ����һ�������������ú���������ִ��
function MyPromise(fn) {
  let _this = this;
  _this.currentState = PENDING;
  _this.value = undefined;
  // ���ڱ��� then �еĻص���ֻ�е� promise
  // ״̬Ϊ pending ʱ�ŻỺ�棬����ÿ��ʵ�����໺��һ��
  _this.resolvedCallbacks = [];
  _this.rejectedCallbacks = [];

  _this.resolve = function (value) {
    if (value instanceof MyPromise) {
      // ��� value �Ǹ� Promise���ݹ�ִ��
      return value.then(_this.resolve, _this.reject)
    }
    setTimeout(() => { // �첽ִ�У���ִ֤��˳��
      if (_this.currentState === PENDING) {
        _this.currentState = RESOLVED;
        _this.value = value;
        _this.resolvedCallbacks.forEach(cb => cb());
      }
    })
  };

  _this.reject = function (reason) {
    setTimeout(() => { // �첽ִ�У���ִ֤��˳��
      if (_this.currentState === PENDING) {
        _this.currentState = REJECTED;
        _this.value = reason;
        _this.rejectedCallbacks.forEach(cb => cb());
      }
    })
  }
  // ���ڽ����������
  // new Promise(() => throw Error('error))
  try {
    fn(_this.resolve, _this.reject);
  } catch (e) {
    _this.reject(e);
  }
}

MyPromise.prototype.then = function (onResolved, onRejected) {
  var self = this;
  // �淶 2.2.7��then ���뷵��һ���µ� promise
  var promise2;
  // �淶 2.2.onResolved �� onRejected ��Ϊ��ѡ����
  // ������Ͳ��Ǻ�����Ҫ���ԣ�ͬʱҲʵ����͸��
  // Promise.resolve(4).then().then((value) => console.log(value))
  onResolved = typeof onResolved === 'function' ? onResolved : v => v;
  onRejected = typeof onRejected === 'function' ? onRejected : r => throw r;

  if (self.currentState === RESOLVED) {
    return (promise2 = new MyPromise(function (resolve, reject) {
      // �淶 2.2.4����֤ onFulfilled��onRjected �첽ִ��
      // �������� setTimeout ������
      setTimeout(function () {
        try {
          var x = onResolved(self.value);
          resolutionProcedure(promise2, x, resolve, reject);
        } catch (reason) {
          reject(reason);
        }
      });
    }));
  }

  if (self.currentState === REJECTED) {
    return (promise2 = new MyPromise(function (resolve, reject) {
      setTimeout(function () {
        // �첽ִ��onRejected
        try {
          var x = onRejected(self.value);
          resolutionProcedure(promise2, x, resolve, reject);
        } catch (reason) {
          reject(reason);
        }
      });
    }));
  }

  if (self.currentState === PENDING) {
    return (promise2 = new MyPromise(function (resolve, reject) {
      self.resolvedCallbacks.push(function () {
        // ���ǵ����ܻ��б�������ʹ�� try/catch ����
        try {
          var x = onResolved(self.value);
          resolutionProcedure(promise2, x, resolve, reject);
        } catch (r) {
          reject(r);
        }
      });

      self.rejectedCallbacks.push(function () {
        try {
          var x = onRejected(self.value);
          resolutionProcedure(promise2, x, resolve, reject);
        } catch (r) {
          reject(r);
        }
      });
    }));
  }
};
// �淶 2.3
function resolutionProcedure(promise2, x, resolve, reject) {
  // �淶 2.3.1��x ���ܺ� promise2 ��ͬ������ѭ������
  if (promise2 === x) {
    return reject(new TypeError("Error"));
  }
  // �淶 2.3.2
  // ��� x Ϊ Promise��״̬Ϊ pending ��Ҫ�����ȴ�����ִ��
  if (x instanceof MyPromise) {
    if (x.currentState === PENDING) {
      x.then(function (value) {
        // �ٴε��øú�����Ϊ��ȷ�� x resolve ��
        // ������ʲô���ͣ�����ǻ������;��ٴ� resolve
        // ��ֵ�����¸� then
        resolutionProcedure(promise2, value, resolve, reject);
      }, reject);
    } else {
      x.then(resolve, reject);
    }
    return;
  }
  // �淶 2.3.3.3.3
  // reject ���� resolve ����һ��ִ�й��û�������������
  let called = false;
  // �淶 2.3.3���ж� x �Ƿ�Ϊ������ߺ���
  if (x !== null && (typeof x === "object" || typeof x === "function")) {
    // �淶 2.3.3.2���������ȡ�� then���� reject
    try {
      // �淶 2.3.3.1
      let then = x.then;
      // ��� then �Ǻ��������� x.then
      if (typeof then === "function") {
        // �淶 2.3.3.3
        then.call(
          x,
          y => {
            if (called) return;
            called = true;
            // �淶 2.3.3.3.1
            resolutionProcedure(promise2, y, resolve, reject);
          },
          e => {
            if (called) return;
            called = true;
            reject(e);
          }
        );
      } else {
        // �淶 2.3.3.4
        resolve(x);
      }
    } catch (e) {
      if (called) return;
      called = true;
      reject(e);
    }
  } else {
    // �淶 2.3.4��x Ϊ��������
    resolve(x);
  }
}

Promise.all = arr => {
    let aResult = [];    //���ڴ��ÿ��ִ�к󷵻ؽ��
    return new _Promise(function (resolve, reject) {
      let i = 0;
      next();    //��ʼ���ִ�������еĺ���
      function next() {
        arr[i].then(function (res) {
          aResult.push(res);    //ִ�к󷵻صĽ������������
          i++;
          if (i == arr.length) {    //������������еĺ�����ִ���꣬��ѽ�����鴫��then
            resolve(aResult);
          } else {
            next();
          }
        })
      }
    })
};