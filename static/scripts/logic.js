$( document ).ready(function() {

    $( "#study" ).show();
    $( "#learning" ).hide();
    $( "#mastered" ).hide();
    
    // Your code here.
    $( "#learning_bt" ).click(function() {
      $( "#learning" ).show();
      $( "#mastered" ).hide();
      $( "#study" ).hide();
    });

    $( "#mastered_bt" ).click(function() {
      $( "#mastered" ).show();
      $( "#learning" ).hide();
      $( "#study" ).hide();
    });

    $( "#study_bt" ).click(function() {
      $( "#study" ).show();
      $( "#learning" ).hide();
      $( "#mastered" ).hide();
    });
});
