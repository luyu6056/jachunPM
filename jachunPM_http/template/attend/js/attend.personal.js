$(document).ready(function()
{
    /* expand active tree. */
    $('.tree li.active .hitarea').click();

    fixHeight();
});

function fixHeight()
{
    var maxHeight = 0;
    $('.main .row .col-xs-4').each(function()
    {
        if(maxHeight < $(this).height()) maxHeight = $(this).height();
    });

    $('.main .row .col-xs-4').height(maxHeight);
}
