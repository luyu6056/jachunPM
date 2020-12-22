$(function()
{
    $('#createForm').ajaxForm({
        beforeSubmit: function(arr, $form, options) {    
            var password = $('#verifyPassword').val()
            var password1 = $('#password1').val();
            var password2 = $('#password2').val();
            var pwd;
            for(var i in arr){
                switch(arr[i].name){
                    case "verifyPassword":
                    pwd = CryptoJS.SHA256(CryptoJS.SHA256(arr[i].value+salt).toString()+rand).toString();
                    arr[i].value=pwd;
                    break;
                }
            }
            for(var i in arr){
                switch(arr[i].name){
                    case "password1":
                    arr[i].value=aes_encrypt(arr[i].value,pwd,'jachunPM');
                    break;
                    case "password2":
                    arr[i].value=aes_encrypt(arr[i].value,pwd,'jachunPM');
                    break;
                }
            }
        },
    })
});