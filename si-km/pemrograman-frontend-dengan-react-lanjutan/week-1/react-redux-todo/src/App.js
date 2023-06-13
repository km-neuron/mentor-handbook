// src/App.js

// import Provider from react-redux
import { Provider } from "react-redux";

// import store

import store from "./redux/store";

// import AddTodoForm component
import AddTodoForm from "./components/AddTodoForm";
// import TodoList component
import TodoList from "./components/TodoList";
//import UserList component
import UserList from "./components/UserList";

function App() {
  return (
    <Provider store={store}>
      <div className="App">
        <UserList />
        <hr />
        <AddTodoForm />
        <hr />
        <TodoList />
      </div>
    </Provider>
  );
}

export default App;