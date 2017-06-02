var app = new Vue({
	el: "#app",
	data: {
		activeform: "",
		result:""
	},
	computed: {
		form: function(){
			if (this.activeform===""){
				return 'bitmap'
			}
			return this.activeform
		}
	}
})
