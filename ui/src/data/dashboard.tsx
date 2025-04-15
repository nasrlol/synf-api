import "./dashboard.css";
import CpuTemperature from "./cpuTempCard";

function dashboard() {
	return (
		<>
			<div className="dashboard">
				<div className="card">
					<h3>CPU Usage</h3>
					<p className="stat-value">{CpuTemperature()}</p>
				</div>
				<div className="card">
					<h3>CPU Temperature</h3>
					<p className="stat-value">0%</p>
				</div>
				<div className="card">
					<h3>RAM Usage</h3>
					<p className="stat-value">0%</p>
				</div>

				<div className="card">
					<h3>Memory Usage</h3>
					<p className="stat-value">0%</p>
				</div>
				<div className="card">
					<h3>Disk Usage</h3>
					<p className="stat-value">0%</p>
				</div>
				<div className="card">
					<h3>Network Status</h3>
					<p className="stat-value">Offline</p>
				</div>
			</div>
		</>
	);
}

export default dashboard;
