<body>

<div id="vue-instance">
    Guess a number between 1 and 10: <input type="number" v-model="userInput">
    <b>{{ message }}</b>
</div>

<script>
    var vm = new Vue({
        el: '#vue-instance',
        data: {
            userInput: 0,
            randomNumber: 5
        },
        methods: {
            getRandomNumber: function(min, max){
                return Math.floor(Math.random() * (max - min + 1)) + min;
            }
        },
        computed: {
            message: function(){
                if (this.userInput == this.randomNumber) {
                    this.randomNumber = this.getRandomNumber(1, 10);
                    return 'You got it right!';
                } else {
                    this.randomNumber = this.getRandomNumber(1, 10);
                    return 'Try again!';
                }

            }
        }
    });
</script>

<script src="http://cdn.jsdelivr.net/vue/1.0.16/vue.js"></script>

</body>
