(function(){
	//˽�о�̬��Ա
	var user = "";
	
	//˽�о�̬����
	function privateStaticMethod(){
	}

	
	Box = function(value){
		//˽�г�Ա
		privateStaticUser = value; 
		
		//�����˽�з���
		function privateMethod(){
		}

		
		//���з�������Ϊ�ܷ���˽�г�Ա��Ҳ����˵����Ȩ������Ҳ����˵��ʵ������
		this.getUser = function(){
			return user;
		};		
		
		//���г�Ա
		this.user = 1;
	};
	
	//���й������
	Box.prototype.sharedMethod = function () {};
	
	//���й�������
	Box.prototype.sharedProperty = function () {};

	
	//���о�̬���� 
	Box.staticMethod = function(){};
	
	//���о�̬��Ա
	Box.staticProperty = 1; 
})();