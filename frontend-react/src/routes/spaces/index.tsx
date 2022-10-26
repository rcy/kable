import { Link } from 'react-router-dom';
import MySpaceMemberships from './MySpaceMemberships';

export default function() {
  return (
    <div className="flex flex-col p-10">
      <section className="pb-10">
        <MySpaceMemberships />
        <Link className="text-blue-600" to="/spaces/explore">explore</Link>
      </section>
    </div>
  );
}