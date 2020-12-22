// Prevent login page show in a iframe modal
if(window.self !== window.top) window.top.location.href = window.location.href;

$(document).ready(function()
{
    /* Fix bug for misc-ping */
    $('#hiddenwin').removeAttr('id');

    var $login = $('#login');
    var adjustPanelPos = function()
    {
        var bestTop = Math.max(0, Math.floor($(window).height() - $login.outerHeight())/2);
        $login.css('margin-top', bestTop);
    };
    adjustPanelPos();
    $(window).on('resize', adjustPanelPos);

    $('#account').focus();

    $("#langs li > a").click(function() 
    {
        selectLang($(this).data('value'));
    });

    $('#loginPanel #submit').click(function()
    {
        var account          = $('#account').val().trim();
        var password         = $('input:password').val().trim();
        var passwordStrength = computePasswordStrength(password);
        var referer   = $('#referer').val();
        var link1     = createLink('user', 'getsalt');
        var link2     = createLink('user', 'login');
        var keepLogin = $('#keepLoginon').attr('checked') == 'checked' ? 1 : 0;
       
        $.ajax
        ({
            url: link1,
            data: 
                {
                    "account": account, 
                },
            dataType: 'json',
            method: 'POST',
            success:function(data)
            {
                if(data.error){
                    alert(data.error);
                    return
                }
                var pwd = CryptoJS.SHA256(CryptoJS.SHA256(password+data.salt).toString()+data.rand).toString();
                $.ajax
                ({
                    url: link2,
                    dataType: 'json',
                    method: 'POST',
                    data: 
                    {
                        "account": account, 
                        "password": pwd,
                        'referer' : referer,
                        'keepLogin' : keepLogin,
                    },
                    success:function(data)
                    {
                        if(data.error) 
                        {
                            alert(data.error);
                            return false;
                        }
                        else
                        {
                            location.href = data.locate;
                        }
                    }
                })
            }
        })
            

    return false;
    })
});
