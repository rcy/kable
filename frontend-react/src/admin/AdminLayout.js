import { Routes, Route, Link } from "react-router-dom";
import PageNotFound from '../PageNotFound';
import AdminAddFamilyMember from './AdminAddFamilyMember.js';
import AdminCreateSpace from './AdminCreateSpace.js';

export default function AdminLayout() {
  return (
    <Routes>
      <Route path="add-family-member" element={<AdminAddFamilyMember/>}/>
      <Route path="create-space" element={<AdminCreateSpace/>}/>
      <Route path="*" element={<PageNotFound />} />
    </Routes>
  )
}

