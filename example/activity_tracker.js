function loadJQuery(callback) {
    const script = document.createElement('script');
    script.src = 'https://code.jquery.com/jquery-3.6.4.min.js';
    script.onload = callback;
    document.head.appendChild(script);
}

loadJQuery(function() {
    const apiBaseUrl = `https://3soiuq2mr0.execute-api.us-east-1.amazonaws.com`;

    function saveActivity(page, data) {
        const apiUrl = `${apiBaseUrl}/activity`;

        const activityData = {
            page: page,
            data: data
        };

        $.ajax({
            url: apiUrl,
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(activityData),
            crossDomain: true,
            success: function(response) {
                console.log('Activity saved successfully:', response);
            },
            error: function(error) {
                console.error('Error saving activity:', error);
            }
        });
    }

    $(document).ready(function() {
        const currentPage = window.location.href;
        const eventData = {
            action: 'page_loaded'
        };

        saveActivity(currentPage, eventData);
    });
});
