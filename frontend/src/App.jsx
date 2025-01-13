import './App.css';

function App() {
    const BACKEND_URL = "http://localhost:8080"
    const fetchTasks = async () => {
        const response = await fetch(BACKEND_URL)
        const data = await response.json()
    }

    return (
        <div className="App">
            
    	</div>
  	);
}

export default App;
