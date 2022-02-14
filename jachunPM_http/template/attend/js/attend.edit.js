$(document).ready(function()
{
    if(status == 'wait')
    {
        $('.editMode').hide();
        $('.viewMode').show();
    }
    else
    {
        $('.editMode').show();
        $('.viewMode').hide();
    }

    $('.edit').click(function()
    {
        $('.editMode').show();
        $('.viewMode').hide();
    })
})
