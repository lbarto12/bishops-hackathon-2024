<script>
    import BarChart from '../lib/charting/BarChart.svelte';
    import PieChart from '../lib/charting/PieChart.svelte';
    import DoughnutChart from '../lib/charting/DoughnutChart.svelte';
</script>

<div class="dashboard">
    <h1>Data Visualization Dashboard</h1>

    <div class="charts">
        <div class="chart">
            <h2>Bar Chart</h2>
            <BarChart />
        </div>

        <div class="chart">
            <h2>Pie Chart</h2>
            <PieChart />
        </div>

        <div class="chart">
            <h2>Doughnut Chart</h2>
            <DoughnutChart />
        </div>
    </div>
</div>

<style>
    .dashboard {
        padding: 20px;
        text-align: center;
        background-color: #f6ecec;
    }

    .charts {
        display: flex;
        justify-content: space-between;
        gap: 20px;
        padding: 20px;
    }

    .chart {
        flex: 1;
        background-color: white;
        padding: 20px;
        border: 1px solid #ddd;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        border-radius: 10px;
    }

    h2 {
        margin-bottom: 20px;
        font-size: 1.5rem;
    }

    @media (max-width: 768px) {
        .charts {
            flex-direction: column;
        }

        .chart {
            margin-bottom: 20px;
        }
    }
</style>




<!-- <script>
    import BarChart from '$lib/charting/BarChart.svelte';
    import PieChart from '$lib/charting/PieChart.svelte';
    import HorizontalBarChart from '$lib/charting/DoughnutChart.svelte';
</script>

<div>
    <h2>Bar Chart</h2>
    <BarChart />

    <PieChart />

    <h2>Horizontal Bar Chart</h2>
    <HorizontalBarChart />
</div>

-->