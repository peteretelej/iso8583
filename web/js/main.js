var app = new Vue({
	el: "#app",
	data: {
		activeform: "bitmap",
		value:"",
		result:"",
	},
	computed: {
		placeholder:function(){
			switch(this.activeform){
				case "bitmap":
					return "Enter Bitmap"
			}
		},
		formtitle: function(){
			switch(this.activeform){
				case "bitmap":
					return "Convert bitmap to binary string"
			}
		}
	},
	methods: {
		bitmapToBinary: function(){
			if (this.value==""){
				newResult("Error: Enter a value into the form for conversion.")
				return
			}
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
					if (error.response.data.response===undefined){
						newResult("Oops! something went wrong. Confirm the value you submitted, must be hex bitmap.")
						return
					}
					newResult("Error: "+ error.response.data.response+".")
				})

		}
	}
})

function newResult(result){
	app.result = result
}
