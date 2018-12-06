(function (window) {
    function Hero(image) {
        this.initialize(image);
    }
    Hero.prototype = new Bitmap();

    Hero.prototype.Bitmap_initialize = Hero.prototype.initialize;
   
    Hero.prototype.initialize = function (image) {
       	this.reset();

        this.Bitmap_initialize(image);
        this.name = 'Hero';
        this.snapToPixel = true;
    };
    Hero.prototype.reset = function() {
    	this.velocity = {x:10*Game.scale,y:25*Game.scale};
       	this.onGround = false;
		this.doubleJump = false;
    };

	Hero.prototype.tick = function () {
		this.velocity.y += 1 * Game.scale;

		// preparing the variables
		var moveBy = {x:0, y:this.velocity.y},
			collision = null,
			collideables = Game.getCollideables();

		collision = calculateCollision(this, 'y', collideables, moveBy);
		// moveBy is now handled by 'calculateCollision'
		// and can also be 0 - therefore we won't have to worry
		this.y += moveBy.y;

		if ( !collision ) {
			if ( this.onGround ) {
				this.onGround = false;
				this.doubleJump = true;
			}
		} else {
			// the hero can only be 'onGround'
			// when he's hitting floor and not
			// some ceiling
			if ( moveBy.y >= 0 ) {
				this.onGround = true;
				this.doubleJump = false;
			}
			this.velocity.y = 0;
		}

		moveBy = {x:this.velocity.x, y:0};
		collision = calculateCollision(this, 'x', collideables, moveBy);
		this.x += moveBy.x;
	};

    Hero.prototype.jump = function() {
    	// if the hero is "on the ground"
    	// let him jump, physically correct!
		if ( this.onGround ) {
			this.velocity.y = -17 * Game.scale;
			this.onGround = false;
			this.doubleJump = true;
		// we want the hero to be able to
		// jump once more when he is in the
		// air - after that, he has to wait
		// to lang somewhere on the ground
		} else if ( this.doubleJump ) {
			this.velocity.y = -17 * Game.scale;
			this.doubleJump = false;
		}
	};

    window.Hero = Hero;
} (window));