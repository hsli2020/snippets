var Register = {  // works, not good
    init: function() {
        Register.print('Register.init');
        Register.validate();
    },
    validate: function() {
        Register.print('Register.validate');
        Register.validateEmail();
        Register.validateUsername();
        Register.validatePassword();
    },
    validateUsername: function() {
        Register.print('Register.validateUsername');
    },
    validateEmail: function() {
        Register.print('Register.validateEmail');
    },
    validatePassword: function() {
        Register.print('Register.validatePassword');
    },
    print: function(string) { // internal function
        console.log(string);
    },
};

Register.init();

Register = (function() {  // better
    function init() {
        print('init');
        validate();
    }
    function validate() {
        print('validate');
        validateEmail();
        validateUsername();
        validatePassword();
    }
    function validateUsername() {
        print('validateUsername');
    }
    function validateEmail() {
        print('validateEmail');
    }
    function validatePassword() {
        print('validatePassword');
    }
    function print(string) { // internal function
        console.log(string);
    }
    return {
        init: init,
    //  validate: validate,
    //  validateUsername: validateUsername,
    //  validateEmail: validateEmail,
    //  validatePassword: validatePassword,
    }
})();

Register.init();
