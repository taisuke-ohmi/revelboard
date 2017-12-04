$(function() {
  $(".revel-form").on("submit", function(event){
    event.preventDefault();
    var $form = $(this);
    var e = $(".revel-form [name=body]");
    var text = e.val();
    text = text.replace(/\r?\n/g, "<br/>");
    e.val(text);
    $.ajax({
      url: "App/PostMessage",
      type: "GET",
      data: $form.serialize(),
      success: function(result, textStatus, xhr) {
        $(".ui.error.message ul").empty();
        if (result && result.length > 0) {
          for (var i = 0; i < result.length; i++) {
            $(".ui.error.message ul").append("<li>"+result[i].Message+"</li>");
          }
          $form[0].reset();
          $(".ui.error.message").css("display", "block");
        } else {
          $form[0].reset();
          $(".ui.error.message").css("display", "none");
        }
      },
      error: function(xhr, textStatus, error) {
        alert("xhr:"+xhr+" status:"+textStatus+" error:"+error);
      }
    });
  });
});
