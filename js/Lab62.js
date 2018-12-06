/**
 * Lab62 - Unique ID generator
 *
 * @license MIT
 * @author Harman Kang <h@h13g.com>
 *
 */
var Lab62 = /** @class */ (function () {
    // Class constructor
    function Lab62() {
        // Init delegate
        this.delegate = {
            b62char: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
            b62string: ""
        };
        // Init generator
        this.generator = function (length) {
            // Run loop specified times
            for (var i = 0; i < length; i++) {
                // Get random integer
                var index = Math.floor(Math.random() * (62));
                // Build ID one character at a time
                this.delegate.b62string += this.delegate.b62char[index];
            }
            return this.delegate.b62string;
        };
    }
    // Make ID
    Lab62.prototype.make = function (length) {
        return this.generator(length);
    };
    return Lab62;
}());
