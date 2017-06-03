var app = new Vue({
	el: "#app",
	data: {
		activeform: "",
		value:"",
		result:"",
	},
	computed: {
		form: function(){
			if (this.activeform===""){
				return 'bitmap'
			}
			return this.activeform
		},
		placeholder:function(){
			switch(this.activeform){
				case "bitmap":
				case "":
					return "Enter Bitmap"
			}
		}

	},
	methods: {
		bitmapToBinary: function(){
			newResult("(fetching result...)")
			axios.get("/api/bitmaptobin",{
				params:{
					msg: this.value
				}
			})
				.then(function(response){
					newResult(response.data.data.Result)
				})
				.catch(function(error){
					console.log(error)
					if (error.data===undefined){
						newResult("Oops! something went wrong. Confirm the value you submitted, must be hex.")
						return
					}
					newResult("Error: "+ error.data.response)
				})

		}
	}
})

function newResult(result){
	app.result = result
}
