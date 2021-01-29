function changeDate(date)
{
    if(date.indexOf('-') != -1)
    {
        var datearray = date.split("-");
        var date = '';
        for(i=0 ; i<datearray.length ; i++)
        {
            date = date + datearray[i]; 
        }
    }
    link = createLink('my', 'effort', 'date=' + date);
    location.href=link;
}
