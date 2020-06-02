var ctx = document.getElementById('myChartDeadLift');
var myChart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: ['January', 'Feburary', 'March', 'April', 'May', 'June'],
        datasets: [{
            data: [10, 14, 18, 16, 14, 19],
            label: "Dead Lift",
            borderColor: "#c45850",
            backgroundColor: "#c45850",
            fill: false
        }],

        borderWidth: 1
    },
    options: {
        title: { display: true, text: "Dead Lift (1RM)" },
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero: false
                }
            }]
        }
    }
});