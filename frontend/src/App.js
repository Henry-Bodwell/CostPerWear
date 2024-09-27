import './App.css';

function App() {
  return (
    <main>
      <h1>Total Wears: <span>256</span></h1>
      <form>
        <div>
          <label>
            Article Name:
            <input
            type="Text" 
            placeholder={'New shirt'}/>
          </label>
          <label>
            Price:
            <input
            type="Float"/>
          </label>

        <label>
          Wears:
          <input
          type="Number"/>
        </label> 

        </div>
        <div>
          <label>
            Tags:
            <select >

            </select>
          </label>
        </div>
        <button type="submit"> Add new Clothing Article</button>
      </form>
    </main>
  );
}

export default App;
