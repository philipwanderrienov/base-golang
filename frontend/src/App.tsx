// src/App.tsx
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Layout from "./components/layout"
import { DashboardPage } from "./pages/dashboard"
import { DataGridPage } from "./pages/data-grid"
import { FormPage } from "./pages/form-page"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<DashboardPage />} />
          <Route path="data" element={<DataGridPage />} />
          <Route path="form" element={<FormPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App