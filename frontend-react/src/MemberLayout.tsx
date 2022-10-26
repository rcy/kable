import { useCurrentFamilyMembershipQuery } from './generated-types';
import { Link, Navigate, NavLink } from 'react-router-dom';
import { Routes, Route } from "react-router-dom";
import MemberHome from './MemberHome';
import AdminLayout from './admin/AdminLayout';
import PageNotFound from './PageNotFound';

import SpacesIndex from './routes/spaces/index';
import SpacesShow from './routes/spaces/show';
import SpacesExplore from './routes/spaces/explore';
import { PersonIdContext } from './contexts'
import { ReactNode } from 'react';

type MemberLayoutType = { doLogout: Function }

export default function MemberLayout({ doLogout }: MemberLayoutType) {
  const { loading, data } = useCurrentFamilyMembershipQuery({ fetchPolicy: 'network-only' })
  if (loading) { return <span>loading</span> }

  return (
    <PersonIdContext.Provider value={data?.currentFamilyMembership?.person?.id}>
      <div className="font-sans">
        <nav className="bg-gray-800 px-2 py-2 text-white flex justify-between text-xl">
          <div className="flex items-center space-x-2">
            <Link to="/">🐙 Octopus Jr</Link>
          </div>
          <div className="text-orange-200">
            <Link to="/me">
              {data?.currentFamilyMembership?.person?.name} ({data?.currentFamilyMembership?.role})
            </Link>
          </div>
        </nav>

        <nav className="bg-black px-10 py-2 text-white flex justify-between text-xl">
          <div>
            <MyNavLink to="/me" activeClass="bg-red-200 text-black">me</MyNavLink>
          </div>
          <div>
            <MyNavLink to="/family" activeClass="bg-blue-200 text-black">my family</MyNavLink>
          </div>
          <div>
            <MyNavLink to="/spaces" activeClass="bg-green-200 text-black">my places</MyNavLink>
          </div>
        </nav>

        <main>
          <Routes>
            <Route path="/" element={<Navigate to="/spaces" />} />
            <Route path="/me" element={<MemberHome familyMembership={data?.currentFamilyMembership} doLogout={doLogout} />} />
            <Route path="/spaces" element={<SpacesIndex />} />
            <Route path="/spaces/explore" element={<SpacesExplore />} />
            <Route path="/spaces/:id" element={<SpacesShow />} />
            <Route path="/admin/*" element={<AdminLayout />} />
            <Route path="*" element={<PageNotFound />} />
          </Routes>
        </main>
      </div>
    </PersonIdContext.Provider>
  )
}

interface MyNavLinkProps {
  to: string,
  activeClass: string,
  children: ReactNode,
}
function MyNavLink({ to, children, activeClass }: MyNavLinkProps) {
  return (
    <NavLink
      to={to}
      className={({ isActive }) => isActive ? activeClass : ""}
    >
      {children}
    </NavLink>
  )
}