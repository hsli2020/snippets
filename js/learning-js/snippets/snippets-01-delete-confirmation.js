//---------------------------------------------------------
// Delete Confirmation (confirmAfterClick)

<a id="delete_link" href="/delete.php">Delete Script</a>

$('#delete_link').click(function(){
    // ֻ��ϵͳ��Ϣ����������, ��������
    return confirm("Are you sure you want to delete?");
    
    // ����javascript���ṩ�ĵ����������첽��ʽ����������
    //bootbox.confirm("Are you sure?", function(result) {
    //    Toast.show("Confirm result: "+result);
    //});
    
    //layer.confirm('Are you sure you want to delete?');

    //return false; // ֻ�м�����һ�䣬javascript���������л�����ʾ
});

// Ϊ��ʹ��javascript��������ֻ�ܲ������ֱ�ͨ�ķ���

<a onclick="mainSmartAdminDelete(this); return false;" 
   data-link="/country/delete/4" 
   href="javascript:void(0);">

function mainSmartAdminDelete(obj) {
   $.SmartMessageBox({
        title   : "Delete",
        content : "Are you sure?",
        buttons : '[No][Yes]'
    }, function(ButtonPressed) {
        if (ButtonPressed === "Yes") {
            redirect($(obj).attr("data-link"));
        }
        if (ButtonPressed === "No") {
            // do nothing
        }
    });
    return false; // ���Ƿ��� false������������������ʾ
}

function redirect(link) {
    window.location = link;
}
