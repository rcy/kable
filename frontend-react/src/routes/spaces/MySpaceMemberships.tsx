import { useContext } from 'react';
import { PersonIdContext } from '../../contexts';
import { useSpaceMembershipsByPersonIdQuery } from "../../generated-types";
import Debug from "../../components/Debug";
import { Link } from 'react-router-dom';

export default function () {
  const personId = useContext(PersonIdContext)

  const queryResult = useSpaceMembershipsByPersonIdQuery({
    variables: {
      personId
    }
  })

  return (
    <div>
      {queryResult.data?.spaceMemberships?.edges.map(({node: m}) => (
        <div key={m.id}>
          <Link to={`/spaces/${m.space?.id}`}>
            {m.space?.name}
          </Link>
        </div>
      ))}
    </div>
  )
}