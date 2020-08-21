https://codesandbox.io/examples/package/bootstrap-vue

<template>		https://codesandbox.io/s/y31zkqnwkz
  <div>
    <h1>Select box</h1>
    <b-dropdown id="ddCommodity"
                  name="ddCommodity"
                  v-model="ddTestVm.ddTestSelectedOption"
                  text="Select Item"
                  variant="primary"
                  class="m-md-2" v-on:change="changeItem">
      <b-dropdown-item disabled value="0">Select an Item</b-dropdown-item>
      <b-dropdown-item v-for="option in ddTestVm.options" 
                        :key="option.value" 
                        :value="option.value"
                        @click="ddTestVm.ddTestSelectedOption = option.value">
        {{option.text}}
      </b-dropdown-item>           
    </b-dropdown> 
    <span>Selected: {{ ddTestVm.ddTestSelectedOption }}</span>
 </div>
</template>

<script>
    export default {
        data() {
            return {
                someOtherProperty: null,
                ddTestVm: {
                    originalValue: [],
                    ddTestSelectedOption: "Value1",
                    disabled: false,
                    readonly: false,
                    visible: true,
                    color: "",
                    options: [
                        {
                            "value": "Value1",
                            "text": "Value1Text"
                        },
                        {
                            "value": "Value2",
                            "text": "Value2Text"
                        },
                        {
                            "value": "Value3",
                            "text": "Value3Text"
                        }
                    ]
                }
            }
        },        
        methods: {
            changeItem: async function () {
            //grab some remote data
                try {
                    let response = await this.$http.get('https://www.example.com/api/' + this.ddTestVm.ddTestSelectedOption + '.json');
                    console.log(response.data);
                    this.someOtherProperty = response.data;
                } catch (error) {
                    console.log(error)
                }
            }
        },
        watch: {

        },
        async created() {

        }
    }
</script>

<style>
</style>