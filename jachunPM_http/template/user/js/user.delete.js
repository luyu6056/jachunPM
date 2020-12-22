$(document).ready(function()
{
     $('#dataform').ajaxForm({
        beforeSubmit: function(arr, $form, options) {    
            for(var i in arr){
                switch(arr[i].name){
                    case "verifyPassword":
                    arr[i].value=CryptoJS.SHA256(CryptoJS.SHA256(arr[i].value+salt).toString()+rand).toString();;
                    break;
                }
            }
        },
    })
});
