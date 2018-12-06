////////////////////////////////////////////////////////////

var Register = {
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
    print: function(string) {
        console.log(string);
    },
};

// Register.init();

////////////////////////////////////////////////////////////

var Register = {
    init: function() {
        this.print('this.init');
        this.validate();
    },
    validate: function() {
        this.print('this.validate');
        this.validateEmail();
        this.validateUsername();
        this.validatePassword();
    },
    validateUsername: function() {
        this.print('this.validateUsername');
    },
    validateEmail: function() {
        this.print('this.validateEmail');
    },
    validatePassword: function() {
        this.print('this.validatePassword');
    },
    print: function(string) {
        console.log(string);
    },
};

// Register.init();

////////////////////////////////////////////////////////////

Register = (function() {
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
    function print(string) {
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

// Register.init();

////////////////////////////////////////////////////////////
