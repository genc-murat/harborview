import { Routes, Route } from "react-router-dom";
import Layout from "./components/layout/Layout";
import Dashboard from "./pages/Dashboard";
import Containers from "./pages/Containers";
import Images from "./pages/Images";
import ContainerDetail from "./pages/ContainerDetail";
import ImageDetail from "./pages/ImageDetail";
import { Toaster } from "@/components/ui/toaster";
import { ToastProvider } from "@/components/ui/use-toast";

function App() {
  return (
    <ToastProvider>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Dashboard />} />
          <Route path="containers" element={<Containers />} />
          <Route path="containers/:id" element={<ContainerDetail />} />
          <Route path="images" element={<Images />} />
          <Route path="images/:id" element={<ImageDetail />} />
        </Route>
      </Routes>
      <Toaster />
    </ToastProvider>
  );
}

export default App;