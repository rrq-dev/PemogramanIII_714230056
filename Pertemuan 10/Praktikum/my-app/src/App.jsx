import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { LayoutDashboard } from "./components/layouts/LayoutDashboard";
import { HomePage } from "./pages/HomePage";
import { MahasiswaPage } from "./pages/MahasiswaPage";
import { TambahMahasiswaPage } from "./pages/TambahMahasiswaPage";
import { Dashboard } from "./pages/Dashboard";
import { EditMahasiswaPage } from "./pages/EditMahasiswaPage";
// import { SettingsPage } from "./pages/SettingsPage";

export default function App() {
  return (
    <Router>
      <LayoutDashboard>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/mahasiswa" element={<MahasiswaPage />} />
          {/*<Route path="/settings" element={<SettingsPage />} />*/}
          <Route path="/" element={<Dashboard />} />
          <Route path="/mahasiswa" element={<MahasiswaPage />} />
          <Route path="/mahasiswa/tambah" element={<TambahMahasiswaPage />} />
          <Route path="/mahasiswa/edit/:npm" element={<EditMahasiswaPage />} />
        </Routes>
      </LayoutDashboard>
    </Router>
  );
}
