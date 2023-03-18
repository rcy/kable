/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  /** A location in a connection that can be used for resuming pagination. */
  Cursor: any;
  /**
   * A point in time as described by the [ISO
   * 8601](https://en.wikipedia.org/wiki/ISO_8601) standard. May or may not include a timezone.
   */
  Datetime: any;
  /** The `JSON` scalar type represents JSON values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf). */
  JSON: any;
  /** A universally unique identifier as defined by [RFC 4122](https://tools.ietf.org/html/rfc4122). */
  UUID: any;
};

export type Authentication = Node & {
  __typename?: 'Authentication';
  createdAt: Scalars['Datetime'];
  details: Scalars['JSON'];
  id: Scalars['UUID'];
  identifier: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  service: Scalars['String'];
  updatedAt: Scalars['Datetime'];
  /** Reads a single `User` that is related to this `Authentication`. */
  user?: Maybe<User>;
  userId: Scalars['UUID'];
};

/**
 * A condition to be used against `Authentication` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type AuthenticationCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `details` field. */
  details?: InputMaybe<Scalars['JSON']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `identifier` field. */
  identifier?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `service` field. */
  service?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `userId` field. */
  userId?: InputMaybe<Scalars['UUID']>;
};

/** A filter to be used against `Authentication` object types. All fields are combined with a logical ‘and.’ */
export type AuthenticationFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<AuthenticationFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `details` field. */
  details?: InputMaybe<JsonFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `identifier` field. */
  identifier?: InputMaybe<StringFilter>;
  /** Negates the expression. */
  not?: InputMaybe<AuthenticationFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<AuthenticationFilter>>;
  /** Filter by the object’s `service` field. */
  service?: InputMaybe<StringFilter>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `user` relation. */
  user?: InputMaybe<UserFilter>;
  /** Filter by the object’s `userId` field. */
  userId?: InputMaybe<UuidFilter>;
};

/** A connection to a list of `Authentication` values. */
export type AuthenticationsConnection = {
  __typename?: 'AuthenticationsConnection';
  /** A list of edges which contains the `Authentication` and cursor to aid in pagination. */
  edges: Array<AuthenticationsEdge>;
  /** A list of `Authentication` objects. */
  nodes: Array<Authentication>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Authentication` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Authentication` edge in the connection. */
export type AuthenticationsEdge = {
  __typename?: 'AuthenticationsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Authentication` at the end of the edge. */
  node: Authentication;
};

/** Methods to use when ordering `Authentication`. */
export enum AuthenticationsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  DetailsAsc = 'DETAILS_ASC',
  DetailsDesc = 'DETAILS_DESC',
  IdentifierAsc = 'IDENTIFIER_ASC',
  IdentifierDesc = 'IDENTIFIER_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  ServiceAsc = 'SERVICE_ASC',
  ServiceDesc = 'SERVICE_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC',
  UserIdAsc = 'USER_ID_ASC',
  UserIdDesc = 'USER_ID_DESC'
}

/** All input for the create `FamilyMembership` mutation. */
export type CreateFamilyMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `FamilyMembership` to be created by this mutation. */
  familyMembership: FamilyMembershipInput;
};

/** The output of our create `FamilyMembership` mutation. */
export type CreateFamilyMembershipPayload = {
  __typename?: 'CreateFamilyMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  /** The `FamilyMembership` that was created by this mutation. */
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our create `FamilyMembership` mutation. */
export type CreateFamilyMembershipPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** All input for the create `Interest` mutation. */
export type CreateInterestInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Interest` to be created by this mutation. */
  interest: InterestInput;
};

/** The output of our create `Interest` mutation. */
export type CreateInterestPayload = {
  __typename?: 'CreateInterestPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Interest` that was created by this mutation. */
  interest?: Maybe<Interest>;
  /** An edge for our `Interest`. May be used by Relay 1. */
  interestEdge?: Maybe<InterestsEdge>;
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
};


/** The output of our create `Interest` mutation. */
export type CreateInterestPayloadInterestEdgeArgs = {
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** All input for the create `ManagedPerson` mutation. */
export type CreateManagedPersonInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `ManagedPerson` to be created by this mutation. */
  managedPerson: ManagedPersonInput;
};

/** The output of our create `ManagedPerson` mutation. */
export type CreateManagedPersonPayload = {
  __typename?: 'CreateManagedPersonPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `ManagedPerson` that was created by this mutation. */
  managedPerson?: Maybe<ManagedPerson>;
  /** An edge for our `ManagedPerson`. May be used by Relay 1. */
  managedPersonEdge?: Maybe<ManagedPeopleEdge>;
  /** Reads a single `Person` that is related to this `ManagedPerson`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `User` that is related to this `ManagedPerson`. */
  user?: Maybe<User>;
};


/** The output of our create `ManagedPerson` mutation. */
export type CreateManagedPersonPayloadManagedPersonEdgeArgs = {
  orderBy?: InputMaybe<Array<ManagedPeopleOrderBy>>;
};

/** All input for the `createNewFamilyMember` mutation. */
export type CreateNewFamilyMemberInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  name: Scalars['String'];
  role: Scalars['String'];
};

/** The output of our `createNewFamilyMember` mutation. */
export type CreateNewFamilyMemberPayload = {
  __typename?: 'CreateNewFamilyMemberPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our `createNewFamilyMember` mutation. */
export type CreateNewFamilyMemberPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** All input for the create `Person` mutation. */
export type CreatePersonInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Person` to be created by this mutation. */
  person: PersonInput;
};

/** The output of our create `Person` mutation. */
export type CreatePersonPayload = {
  __typename?: 'CreatePersonPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Person` that was created by this mutation. */
  person?: Maybe<Person>;
  /** An edge for our `Person`. May be used by Relay 1. */
  personEdge?: Maybe<PeopleEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our create `Person` mutation. */
export type CreatePersonPayloadPersonEdgeArgs = {
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};

/** All input for the create `Space` mutation. */
export type CreateSpaceInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Space` to be created by this mutation. */
  space: SpaceInput;
};

/** All input for the create `SpaceMembership` mutation. */
export type CreateSpaceMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `SpaceMembership` to be created by this mutation. */
  spaceMembership: SpaceMembershipInput;
};

/** The output of our create `SpaceMembership` mutation. */
export type CreateSpaceMembershipPayload = {
  __typename?: 'CreateSpaceMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  /** The `SpaceMembership` that was created by this mutation. */
  spaceMembership?: Maybe<SpaceMembership>;
  /** An edge for our `SpaceMembership`. May be used by Relay 1. */
  spaceMembershipEdge?: Maybe<SpaceMembershipsEdge>;
};


/** The output of our create `SpaceMembership` mutation. */
export type CreateSpaceMembershipPayloadSpaceMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** The output of our create `Space` mutation. */
export type CreateSpacePayload = {
  __typename?: 'CreateSpacePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Space` that was created by this mutation. */
  space?: Maybe<Space>;
  /** An edge for our `Space`. May be used by Relay 1. */
  spaceEdge?: Maybe<SpacesEdge>;
};


/** The output of our create `Space` mutation. */
export type CreateSpacePayloadSpaceEdgeArgs = {
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};

/** All input for the create `Topic` mutation. */
export type CreateTopicInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The `Topic` to be created by this mutation. */
  topic: TopicInput;
};

/** The output of our create `Topic` mutation. */
export type CreateTopicPayload = {
  __typename?: 'CreateTopicPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Topic` that was created by this mutation. */
  topic?: Maybe<Topic>;
  /** An edge for our `Topic`. May be used by Relay 1. */
  topicEdge?: Maybe<TopicsEdge>;
};


/** The output of our create `Topic` mutation. */
export type CreateTopicPayloadTopicEdgeArgs = {
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};

/** All input for the `currentPersonId` mutation. */
export type CurrentPersonIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
};

/** The output of our `currentPersonId` mutation. */
export type CurrentPersonIdPayload = {
  __typename?: 'CurrentPersonIdPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  uuid?: Maybe<Scalars['UUID']>;
};

/** A filter to be used against Datetime fields. All fields are combined with a logical ‘and.’ */
export type DatetimeFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['Datetime']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['Datetime']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['Datetime']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['Datetime']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['Datetime']>>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['Datetime']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['Datetime']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['Datetime']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['Datetime']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['Datetime']>>;
};

/** All input for the `deleteFamilyMembershipByNodeId` mutation. */
export type DeleteFamilyMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `FamilyMembership` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteFamilyMembershipByPersonId` mutation. */
export type DeleteFamilyMembershipByPersonIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  personId: Scalars['UUID'];
};

/** All input for the `deleteFamilyMembership` mutation. */
export type DeleteFamilyMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `FamilyMembership` mutation. */
export type DeleteFamilyMembershipPayload = {
  __typename?: 'DeleteFamilyMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedFamilyMembershipNodeId?: Maybe<Scalars['ID']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  /** The `FamilyMembership` that was deleted by this mutation. */
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our delete `FamilyMembership` mutation. */
export type DeleteFamilyMembershipPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** All input for the `deleteInterestByNodeId` mutation. */
export type DeleteInterestByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Interest` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteInterest` mutation. */
export type DeleteInterestInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `Interest` mutation. */
export type DeleteInterestPayload = {
  __typename?: 'DeleteInterestPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedInterestNodeId?: Maybe<Scalars['ID']>;
  /** The `Interest` that was deleted by this mutation. */
  interest?: Maybe<Interest>;
  /** An edge for our `Interest`. May be used by Relay 1. */
  interestEdge?: Maybe<InterestsEdge>;
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
};


/** The output of our delete `Interest` mutation. */
export type DeleteInterestPayloadInterestEdgeArgs = {
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** All input for the `deletePersonByNodeId` mutation. */
export type DeletePersonByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Person` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deletePerson` mutation. */
export type DeletePersonInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `Person` mutation. */
export type DeletePersonPayload = {
  __typename?: 'DeletePersonPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedPersonNodeId?: Maybe<Scalars['ID']>;
  /** The `Person` that was deleted by this mutation. */
  person?: Maybe<Person>;
  /** An edge for our `Person`. May be used by Relay 1. */
  personEdge?: Maybe<PeopleEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our delete `Person` mutation. */
export type DeletePersonPayloadPersonEdgeArgs = {
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};

/** All input for the `deleteSpaceByNodeId` mutation. */
export type DeleteSpaceByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Space` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteSpace` mutation. */
export type DeleteSpaceInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** All input for the `deleteSpaceMembershipByNodeId` mutation. */
export type DeleteSpaceMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `SpaceMembership` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteSpaceMembershipByPersonIdAndSpaceId` mutation. */
export type DeleteSpaceMembershipByPersonIdAndSpaceIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  personId: Scalars['UUID'];
  spaceId: Scalars['UUID'];
};

/** All input for the `deleteSpaceMembership` mutation. */
export type DeleteSpaceMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `SpaceMembership` mutation. */
export type DeleteSpaceMembershipPayload = {
  __typename?: 'DeleteSpaceMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedSpaceMembershipNodeId?: Maybe<Scalars['ID']>;
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  /** The `SpaceMembership` that was deleted by this mutation. */
  spaceMembership?: Maybe<SpaceMembership>;
  /** An edge for our `SpaceMembership`. May be used by Relay 1. */
  spaceMembershipEdge?: Maybe<SpaceMembershipsEdge>;
};


/** The output of our delete `SpaceMembership` mutation. */
export type DeleteSpaceMembershipPayloadSpaceMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** The output of our delete `Space` mutation. */
export type DeleteSpacePayload = {
  __typename?: 'DeleteSpacePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedSpaceNodeId?: Maybe<Scalars['ID']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Space` that was deleted by this mutation. */
  space?: Maybe<Space>;
  /** An edge for our `Space`. May be used by Relay 1. */
  spaceEdge?: Maybe<SpacesEdge>;
};


/** The output of our delete `Space` mutation. */
export type DeleteSpacePayloadSpaceEdgeArgs = {
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};

/** All input for the `deleteTopicByNodeId` mutation. */
export type DeleteTopicByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Topic` to be deleted. */
  nodeId: Scalars['ID'];
};

/** All input for the `deleteTopic` mutation. */
export type DeleteTopicInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
};

/** The output of our delete `Topic` mutation. */
export type DeleteTopicPayload = {
  __typename?: 'DeleteTopicPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  deletedTopicNodeId?: Maybe<Scalars['ID']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Topic` that was deleted by this mutation. */
  topic?: Maybe<Topic>;
  /** An edge for our `Topic`. May be used by Relay 1. */
  topicEdge?: Maybe<TopicsEdge>;
};


/** The output of our delete `Topic` mutation. */
export type DeleteTopicPayloadTopicEdgeArgs = {
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};

/** A connection to a list of `Family` values. */
export type FamiliesConnection = {
  __typename?: 'FamiliesConnection';
  /** A list of edges which contains the `Family` and cursor to aid in pagination. */
  edges: Array<FamiliesEdge>;
  /** A list of `Family` objects. */
  nodes: Array<Family>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Family` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Family` edge in the connection. */
export type FamiliesEdge = {
  __typename?: 'FamiliesEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Family` at the end of the edge. */
  node: Family;
};

/** Methods to use when ordering `Family`. */
export enum FamiliesOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

export type Family = Node & {
  __typename?: 'Family';
  createdAt: Scalars['Datetime'];
  /** Reads and enables pagination through a set of `FamilyMembership`. */
  familyMemberships: FamilyMembershipsConnection;
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  updatedAt: Scalars['Datetime'];
  /** Reads and enables pagination through a set of `User`. */
  users: UsersConnection;
};


export type FamilyFamilyMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyMembershipCondition>;
  filter?: InputMaybe<FamilyMembershipFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};


export type FamilyUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<UserCondition>;
  filter?: InputMaybe<UserFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<UsersOrderBy>>;
};

/** A condition to be used against `Family` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type FamilyCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `Family` object types. All fields are combined with a logical ‘and.’ */
export type FamilyFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<FamilyFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `familyMemberships` relation. */
  familyMemberships?: InputMaybe<FamilyToManyFamilyMembershipFilter>;
  /** Some related `familyMemberships` exist. */
  familyMembershipsExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Negates the expression. */
  not?: InputMaybe<FamilyFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<FamilyFilter>>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `users` relation. */
  users?: InputMaybe<FamilyToManyUserFilter>;
  /** Some related `users` exist. */
  usersExist?: InputMaybe<Scalars['Boolean']>;
};

export type FamilyMembership = Node & {
  __typename?: 'FamilyMembership';
  createdAt: Scalars['Datetime'];
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  familyId: Scalars['UUID'];
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  personId: Scalars['UUID'];
  role: Scalars['String'];
  title?: Maybe<Scalars['String']>;
};

/**
 * A condition to be used against `FamilyMembership` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type FamilyMembershipCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `familyId` field. */
  familyId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `role` field. */
  role?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `title` field. */
  title?: InputMaybe<Scalars['String']>;
};

/** A filter to be used against `FamilyMembership` object types. All fields are combined with a logical ‘and.’ */
export type FamilyMembershipFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<FamilyMembershipFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `family` relation. */
  family?: InputMaybe<FamilyFilter>;
  /** Filter by the object’s `familyId` field. */
  familyId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Negates the expression. */
  not?: InputMaybe<FamilyMembershipFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<FamilyMembershipFilter>>;
  /** Filter by the object’s `person` relation. */
  person?: InputMaybe<PersonFilter>;
  /** Filter by the object’s `personId` field. */
  personId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `role` field. */
  role?: InputMaybe<StringFilter>;
  /** Filter by the object’s `title` field. */
  title?: InputMaybe<StringFilter>;
};

/** An input for mutations affecting `FamilyMembership` */
export type FamilyMembershipInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  familyId: Scalars['UUID'];
  id?: InputMaybe<Scalars['UUID']>;
  personId: Scalars['UUID'];
  role: Scalars['String'];
  title?: InputMaybe<Scalars['String']>;
};

/** Represents an update to a `FamilyMembership`. Fields that are set will be updated. */
export type FamilyMembershipPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  familyId?: InputMaybe<Scalars['UUID']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  role?: InputMaybe<Scalars['String']>;
  title?: InputMaybe<Scalars['String']>;
};

/** A connection to a list of `FamilyMembership` values. */
export type FamilyMembershipsConnection = {
  __typename?: 'FamilyMembershipsConnection';
  /** A list of edges which contains the `FamilyMembership` and cursor to aid in pagination. */
  edges: Array<FamilyMembershipsEdge>;
  /** A list of `FamilyMembership` objects. */
  nodes: Array<FamilyMembership>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `FamilyMembership` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `FamilyMembership` edge in the connection. */
export type FamilyMembershipsEdge = {
  __typename?: 'FamilyMembershipsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `FamilyMembership` at the end of the edge. */
  node: FamilyMembership;
};

/** Methods to use when ordering `FamilyMembership`. */
export enum FamilyMembershipsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  FamilyIdAsc = 'FAMILY_ID_ASC',
  FamilyIdDesc = 'FAMILY_ID_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  RoleAsc = 'ROLE_ASC',
  RoleDesc = 'ROLE_DESC',
  TitleAsc = 'TITLE_ASC',
  TitleDesc = 'TITLE_DESC'
}

export type FamilyRole = Node & {
  __typename?: 'FamilyRole';
  id: Scalars['Int'];
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
};

/**
 * A condition to be used against `FamilyRole` object types. All fields are tested
 * for equality and combined with a logical ‘and.’
 */
export type FamilyRoleCondition = {
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['Int']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
};

/** A filter to be used against `FamilyRole` object types. All fields are combined with a logical ‘and.’ */
export type FamilyRoleFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<FamilyRoleFilter>>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<IntFilter>;
  /** Filter by the object’s `name` field. */
  name?: InputMaybe<StringFilter>;
  /** Negates the expression. */
  not?: InputMaybe<FamilyRoleFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<FamilyRoleFilter>>;
};

/** A connection to a list of `FamilyRole` values. */
export type FamilyRolesConnection = {
  __typename?: 'FamilyRolesConnection';
  /** A list of edges which contains the `FamilyRole` and cursor to aid in pagination. */
  edges: Array<FamilyRolesEdge>;
  /** A list of `FamilyRole` objects. */
  nodes: Array<FamilyRole>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `FamilyRole` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `FamilyRole` edge in the connection. */
export type FamilyRolesEdge = {
  __typename?: 'FamilyRolesEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `FamilyRole` at the end of the edge. */
  node: FamilyRole;
};

/** Methods to use when ordering `FamilyRole`. */
export enum FamilyRolesOrderBy {
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC'
}

/** A filter to be used against many `FamilyMembership` object types. All fields are combined with a logical ‘and.’ */
export type FamilyToManyFamilyMembershipFilter = {
  /** Every related `FamilyMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<FamilyMembershipFilter>;
  /** No related `FamilyMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<FamilyMembershipFilter>;
  /** Some related `FamilyMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<FamilyMembershipFilter>;
};

/** A filter to be used against many `User` object types. All fields are combined with a logical ‘and.’ */
export type FamilyToManyUserFilter = {
  /** Every related `User` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<UserFilter>;
  /** No related `User` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<UserFilter>;
  /** Some related `User` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<UserFilter>;
};

/** A filter to be used against Int fields. All fields are combined with a logical ‘and.’ */
export type IntFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['Int']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['Int']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['Int']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['Int']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['Int']>>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['Int']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['Int']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['Int']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['Int']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['Int']>>;
};

export type Interest = Node & {
  __typename?: 'Interest';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  personId?: Maybe<Scalars['UUID']>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
  topicId?: Maybe<Scalars['UUID']>;
  updatedAt: Scalars['Datetime'];
};

/**
 * A condition to be used against `Interest` object types. All fields are tested
 * for equality and combined with a logical ‘and.’
 */
export type InterestCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `topicId` field. */
  topicId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `Interest` object types. All fields are combined with a logical ‘and.’ */
export type InterestFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<InterestFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Negates the expression. */
  not?: InputMaybe<InterestFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<InterestFilter>>;
  /** Filter by the object’s `person` relation. */
  person?: InputMaybe<PersonFilter>;
  /** A related `person` exists. */
  personExists?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `personId` field. */
  personId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `topic` relation. */
  topic?: InputMaybe<TopicFilter>;
  /** A related `topic` exists. */
  topicExists?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `topicId` field. */
  topicId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
};

/** An input for mutations affecting `Interest` */
export type InterestInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  topicId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `Interest`. Fields that are set will be updated. */
export type InterestPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  topicId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A connection to a list of `Interest` values. */
export type InterestsConnection = {
  __typename?: 'InterestsConnection';
  /** A list of edges which contains the `Interest` and cursor to aid in pagination. */
  edges: Array<InterestsEdge>;
  /** A list of `Interest` objects. */
  nodes: Array<Interest>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Interest` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Interest` edge in the connection. */
export type InterestsEdge = {
  __typename?: 'InterestsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Interest` at the end of the edge. */
  node: Interest;
};

/** Methods to use when ordering `Interest`. */
export enum InterestsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  TopicIdAsc = 'TOPIC_ID_ASC',
  TopicIdDesc = 'TOPIC_ID_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** A filter to be used against JSON fields. All fields are combined with a logical ‘and.’ */
export type JsonFilter = {
  /** Contained by the specified JSON. */
  containedBy?: InputMaybe<Scalars['JSON']>;
  /** Contains the specified JSON. */
  contains?: InputMaybe<Scalars['JSON']>;
  /** Contains all of the specified keys. */
  containsAllKeys?: InputMaybe<Array<Scalars['String']>>;
  /** Contains any of the specified keys. */
  containsAnyKeys?: InputMaybe<Array<Scalars['String']>>;
  /** Contains the specified key. */
  containsKey?: InputMaybe<Scalars['String']>;
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['JSON']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['JSON']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['JSON']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['JSON']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['JSON']>>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['JSON']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['JSON']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['JSON']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['JSON']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['JSON']>>;
};

export type ListenPayload = {
  __typename?: 'ListenPayload';
  /** Our root query field type. Allows us to run any query from our subscription payload. */
  query?: Maybe<Query>;
  relatedNode?: Maybe<Node>;
  relatedNodeId?: Maybe<Scalars['ID']>;
};

/** A connection to a list of `ManagedPerson` values. */
export type ManagedPeopleConnection = {
  __typename?: 'ManagedPeopleConnection';
  /** A list of edges which contains the `ManagedPerson` and cursor to aid in pagination. */
  edges: Array<ManagedPeopleEdge>;
  /** A list of `ManagedPerson` objects. */
  nodes: Array<ManagedPerson>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `ManagedPerson` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `ManagedPerson` edge in the connection. */
export type ManagedPeopleEdge = {
  __typename?: 'ManagedPeopleEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `ManagedPerson` at the end of the edge. */
  node: ManagedPerson;
};

/** Methods to use when ordering `ManagedPerson`. */
export enum ManagedPeopleOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  UserIdAsc = 'USER_ID_ASC',
  UserIdDesc = 'USER_ID_DESC'
}

export type ManagedPerson = {
  __typename?: 'ManagedPerson';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** Reads a single `Person` that is related to this `ManagedPerson`. */
  person?: Maybe<Person>;
  personId: Scalars['UUID'];
  /** Reads a single `User` that is related to this `ManagedPerson`. */
  user?: Maybe<User>;
  userId: Scalars['UUID'];
};

/**
 * A condition to be used against `ManagedPerson` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type ManagedPersonCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `userId` field. */
  userId?: InputMaybe<Scalars['UUID']>;
};

/** A filter to be used against `ManagedPerson` object types. All fields are combined with a logical ‘and.’ */
export type ManagedPersonFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<ManagedPersonFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Negates the expression. */
  not?: InputMaybe<ManagedPersonFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<ManagedPersonFilter>>;
  /** Filter by the object’s `person` relation. */
  person?: InputMaybe<PersonFilter>;
  /** Filter by the object’s `personId` field. */
  personId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `user` relation. */
  user?: InputMaybe<UserFilter>;
  /** Filter by the object’s `userId` field. */
  userId?: InputMaybe<UuidFilter>;
};

/** An input for mutations affecting `ManagedPerson` */
export type ManagedPersonInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId: Scalars['UUID'];
  userId: Scalars['UUID'];
};

/** The root mutation type which contains root level fields which mutate data. */
export type Mutation = {
  __typename?: 'Mutation';
  /** Creates a single `FamilyMembership`. */
  createFamilyMembership?: Maybe<CreateFamilyMembershipPayload>;
  /** Creates a single `Interest`. */
  createInterest?: Maybe<CreateInterestPayload>;
  /** Creates a single `ManagedPerson`. */
  createManagedPerson?: Maybe<CreateManagedPersonPayload>;
  createNewFamilyMember?: Maybe<CreateNewFamilyMemberPayload>;
  /** Creates a single `Person`. */
  createPerson?: Maybe<CreatePersonPayload>;
  /** Creates a single `Space`. */
  createSpace?: Maybe<CreateSpacePayload>;
  /** Creates a single `SpaceMembership`. */
  createSpaceMembership?: Maybe<CreateSpaceMembershipPayload>;
  /** Creates a single `Topic`. */
  createTopic?: Maybe<CreateTopicPayload>;
  currentPersonId?: Maybe<CurrentPersonIdPayload>;
  /** Deletes a single `FamilyMembership` using a unique key. */
  deleteFamilyMembership?: Maybe<DeleteFamilyMembershipPayload>;
  /** Deletes a single `FamilyMembership` using its globally unique id. */
  deleteFamilyMembershipByNodeId?: Maybe<DeleteFamilyMembershipPayload>;
  /** Deletes a single `FamilyMembership` using a unique key. */
  deleteFamilyMembershipByPersonId?: Maybe<DeleteFamilyMembershipPayload>;
  /** Deletes a single `Interest` using a unique key. */
  deleteInterest?: Maybe<DeleteInterestPayload>;
  /** Deletes a single `Interest` using its globally unique id. */
  deleteInterestByNodeId?: Maybe<DeleteInterestPayload>;
  /** Deletes a single `Person` using a unique key. */
  deletePerson?: Maybe<DeletePersonPayload>;
  /** Deletes a single `Person` using its globally unique id. */
  deletePersonByNodeId?: Maybe<DeletePersonPayload>;
  /** Deletes a single `Space` using a unique key. */
  deleteSpace?: Maybe<DeleteSpacePayload>;
  /** Deletes a single `Space` using its globally unique id. */
  deleteSpaceByNodeId?: Maybe<DeleteSpacePayload>;
  /** Deletes a single `SpaceMembership` using a unique key. */
  deleteSpaceMembership?: Maybe<DeleteSpaceMembershipPayload>;
  /** Deletes a single `SpaceMembership` using its globally unique id. */
  deleteSpaceMembershipByNodeId?: Maybe<DeleteSpaceMembershipPayload>;
  /** Deletes a single `SpaceMembership` using a unique key. */
  deleteSpaceMembershipByPersonIdAndSpaceId?: Maybe<DeleteSpaceMembershipPayload>;
  /** Deletes a single `Topic` using a unique key. */
  deleteTopic?: Maybe<DeleteTopicPayload>;
  /** Deletes a single `Topic` using its globally unique id. */
  deleteTopicByNodeId?: Maybe<DeleteTopicPayload>;
  postMessage?: Maybe<PostMessagePayload>;
  /** Updates a single `FamilyMembership` using a unique key and a patch. */
  updateFamilyMembership?: Maybe<UpdateFamilyMembershipPayload>;
  /** Updates a single `FamilyMembership` using its globally unique id and a patch. */
  updateFamilyMembershipByNodeId?: Maybe<UpdateFamilyMembershipPayload>;
  /** Updates a single `FamilyMembership` using a unique key and a patch. */
  updateFamilyMembershipByPersonId?: Maybe<UpdateFamilyMembershipPayload>;
  /** Updates a single `Interest` using a unique key and a patch. */
  updateInterest?: Maybe<UpdateInterestPayload>;
  /** Updates a single `Interest` using its globally unique id and a patch. */
  updateInterestByNodeId?: Maybe<UpdateInterestPayload>;
  /** Updates a single `Person` using a unique key and a patch. */
  updatePerson?: Maybe<UpdatePersonPayload>;
  /** Updates a single `Person` using its globally unique id and a patch. */
  updatePersonByNodeId?: Maybe<UpdatePersonPayload>;
  /** Updates a single `Space` using a unique key and a patch. */
  updateSpace?: Maybe<UpdateSpacePayload>;
  /** Updates a single `Space` using its globally unique id and a patch. */
  updateSpaceByNodeId?: Maybe<UpdateSpacePayload>;
  /** Updates a single `SpaceMembership` using a unique key and a patch. */
  updateSpaceMembership?: Maybe<UpdateSpaceMembershipPayload>;
  /** Updates a single `SpaceMembership` using its globally unique id and a patch. */
  updateSpaceMembershipByNodeId?: Maybe<UpdateSpaceMembershipPayload>;
  /** Updates a single `SpaceMembership` using a unique key and a patch. */
  updateSpaceMembershipByPersonIdAndSpaceId?: Maybe<UpdateSpaceMembershipPayload>;
  /** Updates a single `Topic` using a unique key and a patch. */
  updateTopic?: Maybe<UpdateTopicPayload>;
  /** Updates a single `Topic` using its globally unique id and a patch. */
  updateTopicByNodeId?: Maybe<UpdateTopicPayload>;
  /** Updates a single `User` using a unique key and a patch. */
  updateUser?: Maybe<UpdateUserPayload>;
  /** Updates a single `User` using its globally unique id and a patch. */
  updateUserByNodeId?: Maybe<UpdateUserPayload>;
  /** Updates a single `User` using a unique key and a patch. */
  updateUserByPersonId?: Maybe<UpdateUserPayload>;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateFamilyMembershipArgs = {
  input: CreateFamilyMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateInterestArgs = {
  input: CreateInterestInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateManagedPersonArgs = {
  input: CreateManagedPersonInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateNewFamilyMemberArgs = {
  input: CreateNewFamilyMemberInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreatePersonArgs = {
  input: CreatePersonInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateSpaceArgs = {
  input: CreateSpaceInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateSpaceMembershipArgs = {
  input: CreateSpaceMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCreateTopicArgs = {
  input: CreateTopicInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationCurrentPersonIdArgs = {
  input: CurrentPersonIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteFamilyMembershipArgs = {
  input: DeleteFamilyMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteFamilyMembershipByNodeIdArgs = {
  input: DeleteFamilyMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteFamilyMembershipByPersonIdArgs = {
  input: DeleteFamilyMembershipByPersonIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteInterestArgs = {
  input: DeleteInterestInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteInterestByNodeIdArgs = {
  input: DeleteInterestByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeletePersonArgs = {
  input: DeletePersonInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeletePersonByNodeIdArgs = {
  input: DeletePersonByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceArgs = {
  input: DeleteSpaceInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceByNodeIdArgs = {
  input: DeleteSpaceByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceMembershipArgs = {
  input: DeleteSpaceMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceMembershipByNodeIdArgs = {
  input: DeleteSpaceMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteSpaceMembershipByPersonIdAndSpaceIdArgs = {
  input: DeleteSpaceMembershipByPersonIdAndSpaceIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteTopicArgs = {
  input: DeleteTopicInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationDeleteTopicByNodeIdArgs = {
  input: DeleteTopicByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationPostMessageArgs = {
  input: PostMessageInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateFamilyMembershipArgs = {
  input: UpdateFamilyMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateFamilyMembershipByNodeIdArgs = {
  input: UpdateFamilyMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateFamilyMembershipByPersonIdArgs = {
  input: UpdateFamilyMembershipByPersonIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateInterestArgs = {
  input: UpdateInterestInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateInterestByNodeIdArgs = {
  input: UpdateInterestByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdatePersonArgs = {
  input: UpdatePersonInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdatePersonByNodeIdArgs = {
  input: UpdatePersonByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceArgs = {
  input: UpdateSpaceInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceByNodeIdArgs = {
  input: UpdateSpaceByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceMembershipArgs = {
  input: UpdateSpaceMembershipInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceMembershipByNodeIdArgs = {
  input: UpdateSpaceMembershipByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateSpaceMembershipByPersonIdAndSpaceIdArgs = {
  input: UpdateSpaceMembershipByPersonIdAndSpaceIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateTopicArgs = {
  input: UpdateTopicInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateTopicByNodeIdArgs = {
  input: UpdateTopicByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateUserArgs = {
  input: UpdateUserInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateUserByNodeIdArgs = {
  input: UpdateUserByNodeIdInput;
};


/** The root mutation type which contains root level fields which mutate data. */
export type MutationUpdateUserByPersonIdArgs = {
  input: UpdateUserByPersonIdInput;
};

/** An object with a globally unique `ID`. */
export type Node = {
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
};

export type Notification = {
  __typename?: 'Notification';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** Reads a single `SpaceMembership` that is related to this `Notification`. */
  membership?: Maybe<SpaceMembership>;
  membershipId: Scalars['UUID'];
  /** Reads a single `Post` that is related to this `Notification`. */
  post?: Maybe<Post>;
  postId: Scalars['UUID'];
};

/**
 * A condition to be used against `Notification` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type NotificationCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `membershipId` field. */
  membershipId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `postId` field. */
  postId?: InputMaybe<Scalars['UUID']>;
};

/** A filter to be used against `Notification` object types. All fields are combined with a logical ‘and.’ */
export type NotificationFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<NotificationFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `membership` relation. */
  membership?: InputMaybe<SpaceMembershipFilter>;
  /** Filter by the object’s `membershipId` field. */
  membershipId?: InputMaybe<UuidFilter>;
  /** Negates the expression. */
  not?: InputMaybe<NotificationFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<NotificationFilter>>;
  /** Filter by the object’s `post` relation. */
  post?: InputMaybe<PostFilter>;
  /** Filter by the object’s `postId` field. */
  postId?: InputMaybe<UuidFilter>;
};

/** A connection to a list of `Notification` values. */
export type NotificationsConnection = {
  __typename?: 'NotificationsConnection';
  /** A list of edges which contains the `Notification` and cursor to aid in pagination. */
  edges: Array<NotificationsEdge>;
  /** A list of `Notification` objects. */
  nodes: Array<Notification>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Notification` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Notification` edge in the connection. */
export type NotificationsEdge = {
  __typename?: 'NotificationsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Notification` at the end of the edge. */
  node: Notification;
};

/** Methods to use when ordering `Notification`. */
export enum NotificationsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  MembershipIdAsc = 'MEMBERSHIP_ID_ASC',
  MembershipIdDesc = 'MEMBERSHIP_ID_DESC',
  Natural = 'NATURAL',
  PostIdAsc = 'POST_ID_ASC',
  PostIdDesc = 'POST_ID_DESC'
}

/** Information about pagination in a connection. */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** When paginating forwards, the cursor to continue. */
  endCursor?: Maybe<Scalars['Cursor']>;
  /** When paginating forwards, are there more items? */
  hasNextPage: Scalars['Boolean'];
  /** When paginating backwards, are there more items? */
  hasPreviousPage: Scalars['Boolean'];
  /** When paginating backwards, the cursor to continue. */
  startCursor?: Maybe<Scalars['Cursor']>;
};

/** A connection to a list of `Person` values. */
export type PeopleConnection = {
  __typename?: 'PeopleConnection';
  /** A list of edges which contains the `Person` and cursor to aid in pagination. */
  edges: Array<PeopleEdge>;
  /** A list of `Person` objects. */
  nodes: Array<Person>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Person` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Person` edge in the connection. */
export type PeopleEdge = {
  __typename?: 'PeopleEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Person` at the end of the edge. */
  node: Person;
};

/** Methods to use when ordering `Person`. */
export enum PeopleOrderBy {
  AvatarUrlAsc = 'AVATAR_URL_ASC',
  AvatarUrlDesc = 'AVATAR_URL_DESC',
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

export type Person = Node & {
  __typename?: 'Person';
  avatarUrl: Scalars['String'];
  createdAt: Scalars['Datetime'];
  /** Reads a single `FamilyMembership` that is related to this `Person`. */
  familyMembership?: Maybe<FamilyMembership>;
  id: Scalars['UUID'];
  /** Reads and enables pagination through a set of `Interest`. */
  interests: InterestsConnection;
  /** Reads and enables pagination through a set of `ManagedPerson`. */
  managedPeople: ManagedPeopleConnection;
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `SpaceMembership`. */
  spaceMemberships: SpaceMembershipsConnection;
  updatedAt: Scalars['Datetime'];
  /** Reads a single `User` that is related to this `Person`. */
  user?: Maybe<User>;
};


export type PersonInterestsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<InterestCondition>;
  filter?: InputMaybe<InterestFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};


export type PersonManagedPeopleArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<ManagedPersonCondition>;
  filter?: InputMaybe<ManagedPersonFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<ManagedPeopleOrderBy>>;
};


export type PersonSpaceMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceMembershipCondition>;
  filter?: InputMaybe<SpaceMembershipFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** A condition to be used against `Person` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type PersonCondition = {
  /** Checks for equality with the object’s `avatarUrl` field. */
  avatarUrl?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `Person` object types. All fields are combined with a logical ‘and.’ */
export type PersonFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<PersonFilter>>;
  /** Filter by the object’s `avatarUrl` field. */
  avatarUrl?: InputMaybe<StringFilter>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `familyMembership` relation. */
  familyMembership?: InputMaybe<FamilyMembershipFilter>;
  /** A related `familyMembership` exists. */
  familyMembershipExists?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `interests` relation. */
  interests?: InputMaybe<PersonToManyInterestFilter>;
  /** Some related `interests` exist. */
  interestsExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `managedPeople` relation. */
  managedPeople?: InputMaybe<PersonToManyManagedPersonFilter>;
  /** Some related `managedPeople` exist. */
  managedPeopleExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `name` field. */
  name?: InputMaybe<StringFilter>;
  /** Negates the expression. */
  not?: InputMaybe<PersonFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<PersonFilter>>;
  /** Filter by the object’s `spaceMemberships` relation. */
  spaceMemberships?: InputMaybe<PersonToManySpaceMembershipFilter>;
  /** Some related `spaceMemberships` exist. */
  spaceMembershipsExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `user` relation. */
  user?: InputMaybe<UserFilter>;
  /** A related `user` exists. */
  userExists?: InputMaybe<Scalars['Boolean']>;
};

/** An input for mutations affecting `Person` */
export type PersonInput = {
  avatarUrl?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name: Scalars['String'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `Person`. Fields that are set will be updated. */
export type PersonPatch = {
  avatarUrl?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against many `Interest` object types. All fields are combined with a logical ‘and.’ */
export type PersonToManyInterestFilter = {
  /** Every related `Interest` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<InterestFilter>;
  /** No related `Interest` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<InterestFilter>;
  /** Some related `Interest` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<InterestFilter>;
};

/** A filter to be used against many `ManagedPerson` object types. All fields are combined with a logical ‘and.’ */
export type PersonToManyManagedPersonFilter = {
  /** Every related `ManagedPerson` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<ManagedPersonFilter>;
  /** No related `ManagedPerson` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<ManagedPersonFilter>;
  /** Some related `ManagedPerson` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<ManagedPersonFilter>;
};

/** A filter to be used against many `SpaceMembership` object types. All fields are combined with a logical ‘and.’ */
export type PersonToManySpaceMembershipFilter = {
  /** Every related `SpaceMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<SpaceMembershipFilter>;
  /** No related `SpaceMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<SpaceMembershipFilter>;
  /** Some related `SpaceMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<SpaceMembershipFilter>;
};

export type Post = Node & {
  __typename?: 'Post';
  body: Scalars['String'];
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** Reads a single `SpaceMembership` that is related to this `Post`. */
  membership?: Maybe<SpaceMembership>;
  membershipId: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `Notification`. */
  notifications: NotificationsConnection;
  /** Reads a single `Space` that is related to this `Post`. */
  space?: Maybe<Space>;
  spaceId: Scalars['UUID'];
  updatedAt: Scalars['Datetime'];
};


export type PostNotificationsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<NotificationCondition>;
  filter?: InputMaybe<NotificationFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<NotificationsOrderBy>>;
};

/** A condition to be used against `Post` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type PostCondition = {
  /** Checks for equality with the object’s `body` field. */
  body?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `membershipId` field. */
  membershipId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `spaceId` field. */
  spaceId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `Post` object types. All fields are combined with a logical ‘and.’ */
export type PostFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<PostFilter>>;
  /** Filter by the object’s `body` field. */
  body?: InputMaybe<StringFilter>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `membership` relation. */
  membership?: InputMaybe<SpaceMembershipFilter>;
  /** Filter by the object’s `membershipId` field. */
  membershipId?: InputMaybe<UuidFilter>;
  /** Negates the expression. */
  not?: InputMaybe<PostFilter>;
  /** Filter by the object’s `notifications` relation. */
  notifications?: InputMaybe<PostToManyNotificationFilter>;
  /** Some related `notifications` exist. */
  notificationsExist?: InputMaybe<Scalars['Boolean']>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<PostFilter>>;
  /** Filter by the object’s `space` relation. */
  space?: InputMaybe<SpaceFilter>;
  /** Filter by the object’s `spaceId` field. */
  spaceId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
};

/** All input for the `postMessage` mutation. */
export type PostMessageInput = {
  body: Scalars['String'];
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  spaceMembershipId: Scalars['UUID'];
};

/** The output of our `postMessage` mutation. */
export type PostMessagePayload = {
  __typename?: 'PostMessagePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `SpaceMembership` that is related to this `Post`. */
  membership?: Maybe<SpaceMembership>;
  post?: Maybe<Post>;
  /** An edge for our `Post`. May be used by Relay 1. */
  postEdge?: Maybe<PostsEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Space` that is related to this `Post`. */
  space?: Maybe<Space>;
};


/** The output of our `postMessage` mutation. */
export type PostMessagePayloadPostEdgeArgs = {
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};

export type PostSubscriptionPayload = {
  __typename?: 'PostSubscriptionPayload';
  event?: Maybe<Scalars['String']>;
  post?: Maybe<Post>;
};

/** A filter to be used against many `Notification` object types. All fields are combined with a logical ‘and.’ */
export type PostToManyNotificationFilter = {
  /** Every related `Notification` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<NotificationFilter>;
  /** No related `Notification` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<NotificationFilter>;
  /** Some related `Notification` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<NotificationFilter>;
};

/** A connection to a list of `Post` values. */
export type PostsConnection = {
  __typename?: 'PostsConnection';
  /** A list of edges which contains the `Post` and cursor to aid in pagination. */
  edges: Array<PostsEdge>;
  /** A list of `Post` objects. */
  nodes: Array<Post>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Post` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Post` edge in the connection. */
export type PostsEdge = {
  __typename?: 'PostsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Post` at the end of the edge. */
  node: Post;
};

/** Methods to use when ordering `Post`. */
export enum PostsOrderBy {
  BodyAsc = 'BODY_ASC',
  BodyDesc = 'BODY_DESC',
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  MembershipIdAsc = 'MEMBERSHIP_ID_ASC',
  MembershipIdDesc = 'MEMBERSHIP_ID_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  SpaceIdAsc = 'SPACE_ID_ASC',
  SpaceIdDesc = 'SPACE_ID_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** The root query type which gives access points into the data universe. */
export type Query = Node & {
  __typename?: 'Query';
  authentication?: Maybe<Authentication>;
  /** Reads a single `Authentication` using its globally unique `ID`. */
  authenticationByNodeId?: Maybe<Authentication>;
  authenticationByServiceAndIdentifier?: Maybe<Authentication>;
  /** Reads and enables pagination through a set of `Authentication`. */
  authentications?: Maybe<AuthenticationsConnection>;
  currentFamilyMembership?: Maybe<FamilyMembership>;
  currentFamilyMembershipId?: Maybe<Scalars['UUID']>;
  currentPerson?: Maybe<Person>;
  currentUser?: Maybe<User>;
  /** Reads and enables pagination through a set of `Family`. */
  families?: Maybe<FamiliesConnection>;
  family?: Maybe<Family>;
  /** Reads a single `Family` using its globally unique `ID`. */
  familyByNodeId?: Maybe<Family>;
  familyMembership?: Maybe<FamilyMembership>;
  /** Reads a single `FamilyMembership` using its globally unique `ID`. */
  familyMembershipByNodeId?: Maybe<FamilyMembership>;
  familyMembershipByPersonId?: Maybe<FamilyMembership>;
  /** Reads and enables pagination through a set of `FamilyMembership`. */
  familyMemberships?: Maybe<FamilyMembershipsConnection>;
  familyRole?: Maybe<FamilyRole>;
  familyRoleByName?: Maybe<FamilyRole>;
  /** Reads a single `FamilyRole` using its globally unique `ID`. */
  familyRoleByNodeId?: Maybe<FamilyRole>;
  /** Reads and enables pagination through a set of `FamilyRole`. */
  familyRoles?: Maybe<FamilyRolesConnection>;
  interest?: Maybe<Interest>;
  /** Reads a single `Interest` using its globally unique `ID`. */
  interestByNodeId?: Maybe<Interest>;
  /** Reads and enables pagination through a set of `Interest`. */
  interests?: Maybe<InterestsConnection>;
  /** Reads and enables pagination through a set of `ManagedPerson`. */
  managedPeople?: Maybe<ManagedPeopleConnection>;
  /** Fetches an object given its globally unique `ID`. */
  node?: Maybe<Node>;
  /** The root query type must be a `Node` to work well with Relay 1 mutations. This just resolves to `query`. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `Notification`. */
  notifications?: Maybe<NotificationsConnection>;
  /** Reads and enables pagination through a set of `Person`. */
  people?: Maybe<PeopleConnection>;
  person?: Maybe<Person>;
  /** Reads a single `Person` using its globally unique `ID`. */
  personByNodeId?: Maybe<Person>;
  post?: Maybe<Post>;
  /** Reads a single `Post` using its globally unique `ID`. */
  postByNodeId?: Maybe<Post>;
  /** Reads and enables pagination through a set of `Post`. */
  posts?: Maybe<PostsConnection>;
  /**
   * Exposes the root query type nested one level down. This is helpful for Relay 1
   * which can only query top level fields if they are in a particular form.
   */
  query: Query;
  space?: Maybe<Space>;
  /** Reads a single `Space` using its globally unique `ID`. */
  spaceByNodeId?: Maybe<Space>;
  spaceMembership?: Maybe<SpaceMembership>;
  /** Reads a single `SpaceMembership` using its globally unique `ID`. */
  spaceMembershipByNodeId?: Maybe<SpaceMembership>;
  spaceMembershipByPersonIdAndSpaceId?: Maybe<SpaceMembership>;
  /** Reads and enables pagination through a set of `SpaceMembership`. */
  spaceMemberships?: Maybe<SpaceMembershipsConnection>;
  /** Reads and enables pagination through a set of `Space`. */
  spaces?: Maybe<SpacesConnection>;
  topic?: Maybe<Topic>;
  /** Reads a single `Topic` using its globally unique `ID`. */
  topicByNodeId?: Maybe<Topic>;
  /** Reads and enables pagination through a set of `Topic`. */
  topics?: Maybe<TopicsConnection>;
  user?: Maybe<User>;
  /** Reads a single `User` using its globally unique `ID`. */
  userByNodeId?: Maybe<User>;
  userByPersonId?: Maybe<User>;
  userId?: Maybe<Scalars['UUID']>;
  /** Reads and enables pagination through a set of `User`. */
  users?: Maybe<UsersConnection>;
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationByServiceAndIdentifierArgs = {
  identifier: Scalars['String'];
  service: Scalars['String'];
};


/** The root query type which gives access points into the data universe. */
export type QueryAuthenticationsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<AuthenticationCondition>;
  filter?: InputMaybe<AuthenticationFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<AuthenticationsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryFamiliesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyCondition>;
  filter?: InputMaybe<FamilyFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamiliesOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyMembershipArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyMembershipByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyMembershipByPersonIdArgs = {
  personId: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyMembershipCondition>;
  filter?: InputMaybe<FamilyMembershipFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRoleArgs = {
  id: Scalars['Int'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRoleByNameArgs = {
  name: Scalars['String'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRoleByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryFamilyRolesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<FamilyRoleCondition>;
  filter?: InputMaybe<FamilyRoleFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<FamilyRolesOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryInterestArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryInterestByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryInterestsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<InterestCondition>;
  filter?: InputMaybe<InterestFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryManagedPeopleArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<ManagedPersonCondition>;
  filter?: InputMaybe<ManagedPersonFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<ManagedPeopleOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryNodeArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryNotificationsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<NotificationCondition>;
  filter?: InputMaybe<NotificationFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<NotificationsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryPeopleArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<PersonCondition>;
  filter?: InputMaybe<PersonFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryPersonArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPersonByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPostArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPostByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryPostsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<PostCondition>;
  filter?: InputMaybe<PostFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceMembershipArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceMembershipByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceMembershipByPersonIdAndSpaceIdArgs = {
  personId: Scalars['UUID'];
  spaceId: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QuerySpaceMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceMembershipCondition>;
  filter?: InputMaybe<SpaceMembershipFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QuerySpacesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceCondition>;
  filter?: InputMaybe<SpaceFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryTopicArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryTopicByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryTopicsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<TopicCondition>;
  filter?: InputMaybe<TopicFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};


/** The root query type which gives access points into the data universe. */
export type QueryUserArgs = {
  id: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryUserByNodeIdArgs = {
  nodeId: Scalars['ID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryUserByPersonIdArgs = {
  personId: Scalars['UUID'];
};


/** The root query type which gives access points into the data universe. */
export type QueryUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<UserCondition>;
  filter?: InputMaybe<UserFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<UsersOrderBy>>;
};

export type Space = Node & {
  __typename?: 'Space';
  createdAt: Scalars['Datetime'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['UUID'];
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `Post`. */
  posts: PostsConnection;
  /** Reads and enables pagination through a set of `SpaceMembership`. */
  spaceMemberships: SpaceMembershipsConnection;
  updatedAt: Scalars['Datetime'];
};


export type SpacePostsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<PostCondition>;
  filter?: InputMaybe<PostFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};


export type SpaceSpaceMembershipsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<SpaceMembershipCondition>;
  filter?: InputMaybe<SpaceMembershipFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** A condition to be used against `Space` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type SpaceCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `description` field. */
  description?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `Space` object types. All fields are combined with a logical ‘and.’ */
export type SpaceFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<SpaceFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `description` field. */
  description?: InputMaybe<StringFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `name` field. */
  name?: InputMaybe<StringFilter>;
  /** Negates the expression. */
  not?: InputMaybe<SpaceFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<SpaceFilter>>;
  /** Filter by the object’s `posts` relation. */
  posts?: InputMaybe<SpaceToManyPostFilter>;
  /** Some related `posts` exist. */
  postsExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `spaceMemberships` relation. */
  spaceMemberships?: InputMaybe<SpaceToManySpaceMembershipFilter>;
  /** Some related `spaceMemberships` exist. */
  spaceMembershipsExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
};

/** An input for mutations affecting `Space` */
export type SpaceInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  description?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['UUID']>;
  name: Scalars['String'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

export type SpaceMembership = Node & {
  __typename?: 'SpaceMembership';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads and enables pagination through a set of `Notification`. */
  notificationsByMembershipId: NotificationsConnection;
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  personId: Scalars['UUID'];
  /** Reads and enables pagination through a set of `Post`. */
  postsByMembershipId: PostsConnection;
  roleId: Scalars['String'];
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  spaceId: Scalars['UUID'];
  updatedAt: Scalars['Datetime'];
};


export type SpaceMembershipNotificationsByMembershipIdArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<NotificationCondition>;
  filter?: InputMaybe<NotificationFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<NotificationsOrderBy>>;
};


export type SpaceMembershipPostsByMembershipIdArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<PostCondition>;
  filter?: InputMaybe<PostFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<PostsOrderBy>>;
};

/**
 * A condition to be used against `SpaceMembership` object types. All fields are
 * tested for equality and combined with a logical ‘and.’
 */
export type SpaceMembershipCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `roleId` field. */
  roleId?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `spaceId` field. */
  spaceId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `SpaceMembership` object types. All fields are combined with a logical ‘and.’ */
export type SpaceMembershipFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<SpaceMembershipFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Negates the expression. */
  not?: InputMaybe<SpaceMembershipFilter>;
  /** Filter by the object’s `notificationsByMembershipId` relation. */
  notificationsByMembershipId?: InputMaybe<SpaceMembershipToManyNotificationFilter>;
  /** Some related `notificationsByMembershipId` exist. */
  notificationsByMembershipIdExist?: InputMaybe<Scalars['Boolean']>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<SpaceMembershipFilter>>;
  /** Filter by the object’s `person` relation. */
  person?: InputMaybe<PersonFilter>;
  /** Filter by the object’s `personId` field. */
  personId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `postsByMembershipId` relation. */
  postsByMembershipId?: InputMaybe<SpaceMembershipToManyPostFilter>;
  /** Some related `postsByMembershipId` exist. */
  postsByMembershipIdExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `roleId` field. */
  roleId?: InputMaybe<StringFilter>;
  /** Filter by the object’s `space` relation. */
  space?: InputMaybe<SpaceFilter>;
  /** Filter by the object’s `spaceId` field. */
  spaceId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
};

/** An input for mutations affecting `SpaceMembership` */
export type SpaceMembershipInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId: Scalars['UUID'];
  roleId: Scalars['String'];
  spaceId: Scalars['UUID'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `SpaceMembership`. Fields that are set will be updated. */
export type SpaceMembershipPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  personId?: InputMaybe<Scalars['UUID']>;
  roleId?: InputMaybe<Scalars['String']>;
  spaceId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against many `Notification` object types. All fields are combined with a logical ‘and.’ */
export type SpaceMembershipToManyNotificationFilter = {
  /** Every related `Notification` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<NotificationFilter>;
  /** No related `Notification` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<NotificationFilter>;
  /** Some related `Notification` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<NotificationFilter>;
};

/** A filter to be used against many `Post` object types. All fields are combined with a logical ‘and.’ */
export type SpaceMembershipToManyPostFilter = {
  /** Every related `Post` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<PostFilter>;
  /** No related `Post` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<PostFilter>;
  /** Some related `Post` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<PostFilter>;
};

/** A connection to a list of `SpaceMembership` values. */
export type SpaceMembershipsConnection = {
  __typename?: 'SpaceMembershipsConnection';
  /** A list of edges which contains the `SpaceMembership` and cursor to aid in pagination. */
  edges: Array<SpaceMembershipsEdge>;
  /** A list of `SpaceMembership` objects. */
  nodes: Array<SpaceMembership>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `SpaceMembership` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `SpaceMembership` edge in the connection. */
export type SpaceMembershipsEdge = {
  __typename?: 'SpaceMembershipsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `SpaceMembership` at the end of the edge. */
  node: SpaceMembership;
};

/** Methods to use when ordering `SpaceMembership`. */
export enum SpaceMembershipsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  RoleIdAsc = 'ROLE_ID_ASC',
  RoleIdDesc = 'ROLE_ID_DESC',
  SpaceIdAsc = 'SPACE_ID_ASC',
  SpaceIdDesc = 'SPACE_ID_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** Represents an update to a `Space`. Fields that are set will be updated. */
export type SpacePatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  description?: InputMaybe<Scalars['String']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against many `Post` object types. All fields are combined with a logical ‘and.’ */
export type SpaceToManyPostFilter = {
  /** Every related `Post` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<PostFilter>;
  /** No related `Post` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<PostFilter>;
  /** Some related `Post` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<PostFilter>;
};

/** A filter to be used against many `SpaceMembership` object types. All fields are combined with a logical ‘and.’ */
export type SpaceToManySpaceMembershipFilter = {
  /** Every related `SpaceMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<SpaceMembershipFilter>;
  /** No related `SpaceMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<SpaceMembershipFilter>;
  /** Some related `SpaceMembership` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<SpaceMembershipFilter>;
};

/** A connection to a list of `Space` values. */
export type SpacesConnection = {
  __typename?: 'SpacesConnection';
  /** A list of edges which contains the `Space` and cursor to aid in pagination. */
  edges: Array<SpacesEdge>;
  /** A list of `Space` objects. */
  nodes: Array<Space>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Space` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Space` edge in the connection. */
export type SpacesEdge = {
  __typename?: 'SpacesEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Space` at the end of the edge. */
  node: Space;
};

/** Methods to use when ordering `Space`. */
export enum SpacesOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  DescriptionAsc = 'DESCRIPTION_ASC',
  DescriptionDesc = 'DESCRIPTION_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** A filter to be used against String fields. All fields are combined with a logical ‘and.’ */
export type StringFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['String']>;
  /** Not equal to the specified value, treating null like an ordinary value (case-insensitive). */
  distinctFromInsensitive?: InputMaybe<Scalars['String']>;
  /** Ends with the specified string (case-sensitive). */
  endsWith?: InputMaybe<Scalars['String']>;
  /** Ends with the specified string (case-insensitive). */
  endsWithInsensitive?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value (case-insensitive). */
  equalToInsensitive?: InputMaybe<Scalars['String']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['String']>;
  /** Greater than the specified value (case-insensitive). */
  greaterThanInsensitive?: InputMaybe<Scalars['String']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['String']>;
  /** Greater than or equal to the specified value (case-insensitive). */
  greaterThanOrEqualToInsensitive?: InputMaybe<Scalars['String']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['String']>>;
  /** Included in the specified list (case-insensitive). */
  inInsensitive?: InputMaybe<Array<Scalars['String']>>;
  /** Contains the specified string (case-sensitive). */
  includes?: InputMaybe<Scalars['String']>;
  /** Contains the specified string (case-insensitive). */
  includesInsensitive?: InputMaybe<Scalars['String']>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['String']>;
  /** Less than the specified value (case-insensitive). */
  lessThanInsensitive?: InputMaybe<Scalars['String']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['String']>;
  /** Less than or equal to the specified value (case-insensitive). */
  lessThanOrEqualToInsensitive?: InputMaybe<Scalars['String']>;
  /** Matches the specified pattern (case-sensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  like?: InputMaybe<Scalars['String']>;
  /** Matches the specified pattern (case-insensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  likeInsensitive?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['String']>;
  /** Equal to the specified value, treating null like an ordinary value (case-insensitive). */
  notDistinctFromInsensitive?: InputMaybe<Scalars['String']>;
  /** Does not end with the specified string (case-sensitive). */
  notEndsWith?: InputMaybe<Scalars['String']>;
  /** Does not end with the specified string (case-insensitive). */
  notEndsWithInsensitive?: InputMaybe<Scalars['String']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['String']>;
  /** Not equal to the specified value (case-insensitive). */
  notEqualToInsensitive?: InputMaybe<Scalars['String']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['String']>>;
  /** Not included in the specified list (case-insensitive). */
  notInInsensitive?: InputMaybe<Array<Scalars['String']>>;
  /** Does not contain the specified string (case-sensitive). */
  notIncludes?: InputMaybe<Scalars['String']>;
  /** Does not contain the specified string (case-insensitive). */
  notIncludesInsensitive?: InputMaybe<Scalars['String']>;
  /** Does not match the specified pattern (case-sensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  notLike?: InputMaybe<Scalars['String']>;
  /** Does not match the specified pattern (case-insensitive). An underscore (_) matches any single character; a percent sign (%) matches any sequence of zero or more characters. */
  notLikeInsensitive?: InputMaybe<Scalars['String']>;
  /** Does not start with the specified string (case-sensitive). */
  notStartsWith?: InputMaybe<Scalars['String']>;
  /** Does not start with the specified string (case-insensitive). */
  notStartsWithInsensitive?: InputMaybe<Scalars['String']>;
  /** Starts with the specified string (case-sensitive). */
  startsWith?: InputMaybe<Scalars['String']>;
  /** Starts with the specified string (case-insensitive). */
  startsWithInsensitive?: InputMaybe<Scalars['String']>;
};

/** The root subscription type: contains realtime events you can subscribe to with the `subscription` operation. */
export type Subscription = {
  __typename?: 'Subscription';
  listen: ListenPayload;
  posts?: Maybe<PostSubscriptionPayload>;
};


/** The root subscription type: contains realtime events you can subscribe to with the `subscription` operation. */
export type SubscriptionListenArgs = {
  topic: Scalars['String'];
};


/** The root subscription type: contains realtime events you can subscribe to with the `subscription` operation. */
export type SubscriptionPostsArgs = {
  spaceId: Scalars['UUID'];
};

export type Topic = Node & {
  __typename?: 'Topic';
  createdAt: Scalars['Datetime'];
  id: Scalars['UUID'];
  /** Reads and enables pagination through a set of `Interest`. */
  interests: InterestsConnection;
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  updatedAt: Scalars['Datetime'];
};


export type TopicInterestsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<InterestCondition>;
  filter?: InputMaybe<InterestFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** A condition to be used against `Topic` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type TopicCondition = {
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `Topic` object types. All fields are combined with a logical ‘and.’ */
export type TopicFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<TopicFilter>>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `interests` relation. */
  interests?: InputMaybe<TopicToManyInterestFilter>;
  /** Some related `interests` exist. */
  interestsExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `name` field. */
  name?: InputMaybe<StringFilter>;
  /** Negates the expression. */
  not?: InputMaybe<TopicFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<TopicFilter>>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
};

/** An input for mutations affecting `Topic` */
export type TopicInput = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name: Scalars['String'];
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** Represents an update to a `Topic`. Fields that are set will be updated. */
export type TopicPatch = {
  createdAt?: InputMaybe<Scalars['Datetime']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against many `Interest` object types. All fields are combined with a logical ‘and.’ */
export type TopicToManyInterestFilter = {
  /** Every related `Interest` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<InterestFilter>;
  /** No related `Interest` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<InterestFilter>;
  /** Some related `Interest` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<InterestFilter>;
};

/** A connection to a list of `Topic` values. */
export type TopicsConnection = {
  __typename?: 'TopicsConnection';
  /** A list of edges which contains the `Topic` and cursor to aid in pagination. */
  edges: Array<TopicsEdge>;
  /** A list of `Topic` objects. */
  nodes: Array<Topic>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `Topic` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `Topic` edge in the connection. */
export type TopicsEdge = {
  __typename?: 'TopicsEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `Topic` at the end of the edge. */
  node: Topic;
};

/** Methods to use when ordering `Topic`. */
export enum TopicsOrderBy {
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

/** A filter to be used against UUID fields. All fields are combined with a logical ‘and.’ */
export type UuidFilter = {
  /** Not equal to the specified value, treating null like an ordinary value. */
  distinctFrom?: InputMaybe<Scalars['UUID']>;
  /** Equal to the specified value. */
  equalTo?: InputMaybe<Scalars['UUID']>;
  /** Greater than the specified value. */
  greaterThan?: InputMaybe<Scalars['UUID']>;
  /** Greater than or equal to the specified value. */
  greaterThanOrEqualTo?: InputMaybe<Scalars['UUID']>;
  /** Included in the specified list. */
  in?: InputMaybe<Array<Scalars['UUID']>>;
  /** Is null (if `true` is specified) or is not null (if `false` is specified). */
  isNull?: InputMaybe<Scalars['Boolean']>;
  /** Less than the specified value. */
  lessThan?: InputMaybe<Scalars['UUID']>;
  /** Less than or equal to the specified value. */
  lessThanOrEqualTo?: InputMaybe<Scalars['UUID']>;
  /** Equal to the specified value, treating null like an ordinary value. */
  notDistinctFrom?: InputMaybe<Scalars['UUID']>;
  /** Not equal to the specified value. */
  notEqualTo?: InputMaybe<Scalars['UUID']>;
  /** Not included in the specified list. */
  notIn?: InputMaybe<Array<Scalars['UUID']>>;
};

/** All input for the `updateFamilyMembershipByNodeId` mutation. */
export type UpdateFamilyMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `FamilyMembership` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `FamilyMembership` being updated. */
  patch: FamilyMembershipPatch;
};

/** All input for the `updateFamilyMembershipByPersonId` mutation. */
export type UpdateFamilyMembershipByPersonIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** An object where the defined keys will be set on the `FamilyMembership` being updated. */
  patch: FamilyMembershipPatch;
  personId: Scalars['UUID'];
};

/** All input for the `updateFamilyMembership` mutation. */
export type UpdateFamilyMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `FamilyMembership` being updated. */
  patch: FamilyMembershipPatch;
};

/** The output of our update `FamilyMembership` mutation. */
export type UpdateFamilyMembershipPayload = {
  __typename?: 'UpdateFamilyMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Family` that is related to this `FamilyMembership`. */
  family?: Maybe<Family>;
  /** The `FamilyMembership` that was updated by this mutation. */
  familyMembership?: Maybe<FamilyMembership>;
  /** An edge for our `FamilyMembership`. May be used by Relay 1. */
  familyMembershipEdge?: Maybe<FamilyMembershipsEdge>;
  /** Reads a single `Person` that is related to this `FamilyMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our update `FamilyMembership` mutation. */
export type UpdateFamilyMembershipPayloadFamilyMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<FamilyMembershipsOrderBy>>;
};

/** All input for the `updateInterestByNodeId` mutation. */
export type UpdateInterestByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Interest` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Interest` being updated. */
  patch: InterestPatch;
};

/** All input for the `updateInterest` mutation. */
export type UpdateInterestInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Interest` being updated. */
  patch: InterestPatch;
};

/** The output of our update `Interest` mutation. */
export type UpdateInterestPayload = {
  __typename?: 'UpdateInterestPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Interest` that was updated by this mutation. */
  interest?: Maybe<Interest>;
  /** An edge for our `Interest`. May be used by Relay 1. */
  interestEdge?: Maybe<InterestsEdge>;
  /** Reads a single `Person` that is related to this `Interest`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Topic` that is related to this `Interest`. */
  topic?: Maybe<Topic>;
};


/** The output of our update `Interest` mutation. */
export type UpdateInterestPayloadInterestEdgeArgs = {
  orderBy?: InputMaybe<Array<InterestsOrderBy>>;
};

/** All input for the `updatePersonByNodeId` mutation. */
export type UpdatePersonByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Person` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Person` being updated. */
  patch: PersonPatch;
};

/** All input for the `updatePerson` mutation. */
export type UpdatePersonInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Person` being updated. */
  patch: PersonPatch;
};

/** The output of our update `Person` mutation. */
export type UpdatePersonPayload = {
  __typename?: 'UpdatePersonPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** The `Person` that was updated by this mutation. */
  person?: Maybe<Person>;
  /** An edge for our `Person`. May be used by Relay 1. */
  personEdge?: Maybe<PeopleEdge>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
};


/** The output of our update `Person` mutation. */
export type UpdatePersonPayloadPersonEdgeArgs = {
  orderBy?: InputMaybe<Array<PeopleOrderBy>>;
};

/** All input for the `updateSpaceByNodeId` mutation. */
export type UpdateSpaceByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Space` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Space` being updated. */
  patch: SpacePatch;
};

/** All input for the `updateSpace` mutation. */
export type UpdateSpaceInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Space` being updated. */
  patch: SpacePatch;
};

/** All input for the `updateSpaceMembershipByNodeId` mutation. */
export type UpdateSpaceMembershipByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `SpaceMembership` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `SpaceMembership` being updated. */
  patch: SpaceMembershipPatch;
};

/** All input for the `updateSpaceMembershipByPersonIdAndSpaceId` mutation. */
export type UpdateSpaceMembershipByPersonIdAndSpaceIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** An object where the defined keys will be set on the `SpaceMembership` being updated. */
  patch: SpaceMembershipPatch;
  personId: Scalars['UUID'];
  spaceId: Scalars['UUID'];
};

/** All input for the `updateSpaceMembership` mutation. */
export type UpdateSpaceMembershipInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `SpaceMembership` being updated. */
  patch: SpaceMembershipPatch;
};

/** The output of our update `SpaceMembership` mutation. */
export type UpdateSpaceMembershipPayload = {
  __typename?: 'UpdateSpaceMembershipPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Person` that is related to this `SpaceMembership`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** Reads a single `Space` that is related to this `SpaceMembership`. */
  space?: Maybe<Space>;
  /** The `SpaceMembership` that was updated by this mutation. */
  spaceMembership?: Maybe<SpaceMembership>;
  /** An edge for our `SpaceMembership`. May be used by Relay 1. */
  spaceMembershipEdge?: Maybe<SpaceMembershipsEdge>;
};


/** The output of our update `SpaceMembership` mutation. */
export type UpdateSpaceMembershipPayloadSpaceMembershipEdgeArgs = {
  orderBy?: InputMaybe<Array<SpaceMembershipsOrderBy>>;
};

/** The output of our update `Space` mutation. */
export type UpdateSpacePayload = {
  __typename?: 'UpdateSpacePayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Space` that was updated by this mutation. */
  space?: Maybe<Space>;
  /** An edge for our `Space`. May be used by Relay 1. */
  spaceEdge?: Maybe<SpacesEdge>;
};


/** The output of our update `Space` mutation. */
export type UpdateSpacePayloadSpaceEdgeArgs = {
  orderBy?: InputMaybe<Array<SpacesOrderBy>>;
};

/** All input for the `updateTopicByNodeId` mutation. */
export type UpdateTopicByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `Topic` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `Topic` being updated. */
  patch: TopicPatch;
};

/** All input for the `updateTopic` mutation. */
export type UpdateTopicInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `Topic` being updated. */
  patch: TopicPatch;
};

/** The output of our update `Topic` mutation. */
export type UpdateTopicPayload = {
  __typename?: 'UpdateTopicPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `Topic` that was updated by this mutation. */
  topic?: Maybe<Topic>;
  /** An edge for our `Topic`. May be used by Relay 1. */
  topicEdge?: Maybe<TopicsEdge>;
};


/** The output of our update `Topic` mutation. */
export type UpdateTopicPayloadTopicEdgeArgs = {
  orderBy?: InputMaybe<Array<TopicsOrderBy>>;
};

/** All input for the `updateUserByNodeId` mutation. */
export type UpdateUserByNodeIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** The globally unique `ID` which will identify a single `User` to be updated. */
  nodeId: Scalars['ID'];
  /** An object where the defined keys will be set on the `User` being updated. */
  patch: UserPatch;
};

/** All input for the `updateUserByPersonId` mutation. */
export type UpdateUserByPersonIdInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  /** An object where the defined keys will be set on the `User` being updated. */
  patch: UserPatch;
  personId: Scalars['UUID'];
};

/** All input for the `updateUser` mutation. */
export type UpdateUserInput = {
  /**
   * An arbitrary string value with no semantic meaning. Will be included in the
   * payload verbatim. May be used to track mutations by the client.
   */
  clientMutationId?: InputMaybe<Scalars['String']>;
  id: Scalars['UUID'];
  /** An object where the defined keys will be set on the `User` being updated. */
  patch: UserPatch;
};

/** The output of our update `User` mutation. */
export type UpdateUserPayload = {
  __typename?: 'UpdateUserPayload';
  /**
   * The exact same `clientMutationId` that was provided in the mutation input,
   * unchanged and unused. May be used by a client to track mutations.
   */
  clientMutationId?: Maybe<Scalars['String']>;
  /** Reads a single `Family` that is related to this `User`. */
  family?: Maybe<Family>;
  /** Reads a single `Person` that is related to this `User`. */
  person?: Maybe<Person>;
  /** Our root query field type. Allows us to run any query from our mutation payload. */
  query?: Maybe<Query>;
  /** The `User` that was updated by this mutation. */
  user?: Maybe<User>;
  /** An edge for our `User`. May be used by Relay 1. */
  userEdge?: Maybe<UsersEdge>;
};


/** The output of our update `User` mutation. */
export type UpdateUserPayloadUserEdgeArgs = {
  orderBy?: InputMaybe<Array<UsersOrderBy>>;
};

export type User = Node & {
  __typename?: 'User';
  /** Reads and enables pagination through a set of `Authentication`. */
  authentications: AuthenticationsConnection;
  avatarUrl?: Maybe<Scalars['String']>;
  createdAt: Scalars['Datetime'];
  /** Reads a single `Family` that is related to this `User`. */
  family?: Maybe<Family>;
  familyId?: Maybe<Scalars['UUID']>;
  id: Scalars['UUID'];
  /** Reads and enables pagination through a set of `ManagedPerson`. */
  managedPeople: ManagedPeopleConnection;
  name: Scalars['String'];
  /** A globally unique identifier. Can be used in various places throughout the system to identify this single value. */
  nodeId: Scalars['ID'];
  /** Reads a single `Person` that is related to this `User`. */
  person?: Maybe<Person>;
  personId?: Maybe<Scalars['UUID']>;
  updatedAt: Scalars['Datetime'];
};


export type UserAuthenticationsArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<AuthenticationCondition>;
  filter?: InputMaybe<AuthenticationFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<AuthenticationsOrderBy>>;
};


export type UserManagedPeopleArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  condition?: InputMaybe<ManagedPersonCondition>;
  filter?: InputMaybe<ManagedPersonFilter>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<Array<ManagedPeopleOrderBy>>;
};

/** A condition to be used against `User` object types. All fields are tested for equality and combined with a logical ‘and.’ */
export type UserCondition = {
  /** Checks for equality with the object’s `avatarUrl` field. */
  avatarUrl?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `createdAt` field. */
  createdAt?: InputMaybe<Scalars['Datetime']>;
  /** Checks for equality with the object’s `familyId` field. */
  familyId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `id` field. */
  id?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `name` field. */
  name?: InputMaybe<Scalars['String']>;
  /** Checks for equality with the object’s `personId` field. */
  personId?: InputMaybe<Scalars['UUID']>;
  /** Checks for equality with the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against `User` object types. All fields are combined with a logical ‘and.’ */
export type UserFilter = {
  /** Checks for all expressions in this list. */
  and?: InputMaybe<Array<UserFilter>>;
  /** Filter by the object’s `authentications` relation. */
  authentications?: InputMaybe<UserToManyAuthenticationFilter>;
  /** Some related `authentications` exist. */
  authenticationsExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `avatarUrl` field. */
  avatarUrl?: InputMaybe<StringFilter>;
  /** Filter by the object’s `createdAt` field. */
  createdAt?: InputMaybe<DatetimeFilter>;
  /** Filter by the object’s `family` relation. */
  family?: InputMaybe<FamilyFilter>;
  /** A related `family` exists. */
  familyExists?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `familyId` field. */
  familyId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `id` field. */
  id?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `managedPeople` relation. */
  managedPeople?: InputMaybe<UserToManyManagedPersonFilter>;
  /** Some related `managedPeople` exist. */
  managedPeopleExist?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `name` field. */
  name?: InputMaybe<StringFilter>;
  /** Negates the expression. */
  not?: InputMaybe<UserFilter>;
  /** Checks for any expressions in this list. */
  or?: InputMaybe<Array<UserFilter>>;
  /** Filter by the object’s `person` relation. */
  person?: InputMaybe<PersonFilter>;
  /** A related `person` exists. */
  personExists?: InputMaybe<Scalars['Boolean']>;
  /** Filter by the object’s `personId` field. */
  personId?: InputMaybe<UuidFilter>;
  /** Filter by the object’s `updatedAt` field. */
  updatedAt?: InputMaybe<DatetimeFilter>;
};

/** Represents an update to a `User`. Fields that are set will be updated. */
export type UserPatch = {
  avatarUrl?: InputMaybe<Scalars['String']>;
  createdAt?: InputMaybe<Scalars['Datetime']>;
  familyId?: InputMaybe<Scalars['UUID']>;
  id?: InputMaybe<Scalars['UUID']>;
  name?: InputMaybe<Scalars['String']>;
  personId?: InputMaybe<Scalars['UUID']>;
  updatedAt?: InputMaybe<Scalars['Datetime']>;
};

/** A filter to be used against many `Authentication` object types. All fields are combined with a logical ‘and.’ */
export type UserToManyAuthenticationFilter = {
  /** Every related `Authentication` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<AuthenticationFilter>;
  /** No related `Authentication` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<AuthenticationFilter>;
  /** Some related `Authentication` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<AuthenticationFilter>;
};

/** A filter to be used against many `ManagedPerson` object types. All fields are combined with a logical ‘and.’ */
export type UserToManyManagedPersonFilter = {
  /** Every related `ManagedPerson` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  every?: InputMaybe<ManagedPersonFilter>;
  /** No related `ManagedPerson` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  none?: InputMaybe<ManagedPersonFilter>;
  /** Some related `ManagedPerson` matches the filter criteria. All fields are combined with a logical ‘and.’ */
  some?: InputMaybe<ManagedPersonFilter>;
};

/** A connection to a list of `User` values. */
export type UsersConnection = {
  __typename?: 'UsersConnection';
  /** A list of edges which contains the `User` and cursor to aid in pagination. */
  edges: Array<UsersEdge>;
  /** A list of `User` objects. */
  nodes: Array<User>;
  /** Information to aid in pagination. */
  pageInfo: PageInfo;
  /** The count of *all* `User` you could get from the connection. */
  totalCount: Scalars['Int'];
};

/** A `User` edge in the connection. */
export type UsersEdge = {
  __typename?: 'UsersEdge';
  /** A cursor for use in pagination. */
  cursor?: Maybe<Scalars['Cursor']>;
  /** The `User` at the end of the edge. */
  node: User;
};

/** Methods to use when ordering `User`. */
export enum UsersOrderBy {
  AvatarUrlAsc = 'AVATAR_URL_ASC',
  AvatarUrlDesc = 'AVATAR_URL_DESC',
  CreatedAtAsc = 'CREATED_AT_ASC',
  CreatedAtDesc = 'CREATED_AT_DESC',
  FamilyIdAsc = 'FAMILY_ID_ASC',
  FamilyIdDesc = 'FAMILY_ID_DESC',
  IdAsc = 'ID_ASC',
  IdDesc = 'ID_DESC',
  NameAsc = 'NAME_ASC',
  NameDesc = 'NAME_DESC',
  Natural = 'NATURAL',
  PersonIdAsc = 'PERSON_ID_ASC',
  PersonIdDesc = 'PERSON_ID_DESC',
  PrimaryKeyAsc = 'PRIMARY_KEY_ASC',
  PrimaryKeyDesc = 'PRIMARY_KEY_DESC',
  UpdatedAtAsc = 'UPDATED_AT_ASC',
  UpdatedAtDesc = 'UPDATED_AT_DESC'
}

export type AllSpacesQueryVariables = Exact<{
  personId: Scalars['UUID'];
}>;


export type AllSpacesQuery = { __typename?: 'Query', spaces?: { __typename?: 'SpacesConnection', totalCount: number, edges: Array<{ __typename?: 'SpacesEdge', node: { __typename?: 'Space', id: any, name: string, spaceMemberships: { __typename?: 'SpaceMembershipsConnection', edges: Array<{ __typename?: 'SpaceMembershipsEdge', node: { __typename?: 'SpaceMembership', id: any, person?: { __typename?: 'Person', id: any } | null } }> } } }> } | null };

export type CreateNewFamilyMemberMutationVariables = Exact<{
  name: Scalars['String'];
  role: Scalars['String'];
}>;


export type CreateNewFamilyMemberMutation = { __typename?: 'Mutation', createNewFamilyMember?: { __typename?: 'CreateNewFamilyMemberPayload', familyMembership?: { __typename?: 'FamilyMembership', personId: any } | null } | null };

export type CreateSpaceMutationVariables = Exact<{
  name: Scalars['String'];
}>;


export type CreateSpaceMutation = { __typename?: 'Mutation', createSpace?: { __typename?: 'CreateSpacePayload', space?: { __typename?: 'Space', id: any } | null } | null };

export type CurrentPersonQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentPersonQuery = { __typename?: 'Query', currentPerson?: { __typename?: 'Person', id: any, name: string, avatarUrl: string } | null };

export type CurrentUserQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserQuery = { __typename?: 'Query', currentUser?: { __typename?: 'User', id: any, name: string } | null };

export type CurrentUserFamilyQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserFamilyQuery = { __typename?: 'Query', currentUser?: { __typename?: 'User', id: any, name: string, family?: { __typename?: 'Family', id: any, familyMemberships: { __typename?: 'FamilyMembershipsConnection', nodes: Array<{ __typename?: 'FamilyMembership', id: any, title?: string | null, role: string, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string, user?: { __typename?: 'User', id: any } | null } | null }> } } | null } | null };

export type CurrentFamilyMembershipQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentFamilyMembershipQuery = { __typename?: 'Query', currentFamilyMembership?: { __typename?: 'FamilyMembership', id: any, role: string, title?: string | null, family?: { __typename?: 'Family', id: any } | null, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string, user?: { __typename?: 'User', id: any } | null } | null } | null };

export type CurrentUserWithManagedPeopleQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserWithManagedPeopleQuery = { __typename?: 'Query', currentUser?: { __typename?: 'User', id: any, name: string, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string } | null, managedPeople: { __typename?: 'ManagedPeopleConnection', nodes: Array<{ __typename?: 'ManagedPerson', id: any, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string } | null }> } } | null };

export type JoinSpaceMutationVariables = Exact<{
  spaceId: Scalars['UUID'];
  personId: Scalars['UUID'];
}>;


export type JoinSpaceMutation = { __typename?: 'Mutation', createSpaceMembership?: { __typename?: 'CreateSpaceMembershipPayload', clientMutationId?: string | null } | null };

export type PostMessageMutationVariables = Exact<{
  membershipId: Scalars['UUID'];
  body: Scalars['String'];
}>;


export type PostMessageMutation = { __typename?: 'Mutation', postMessage?: { __typename?: 'PostMessagePayload', post?: { __typename?: 'Post', id: any } | null } | null };

export type SetPersonAvatarMutationVariables = Exact<{
  personId: Scalars['UUID'];
  avatarUrl: Scalars['String'];
}>;


export type SetPersonAvatarMutation = { __typename?: 'Mutation', updatePerson?: { __typename?: 'UpdatePersonPayload', clientMutationId?: string | null } | null };

export type SpaceQueryVariables = Exact<{
  id: Scalars['UUID'];
}>;


export type SpaceQuery = { __typename?: 'Query', space?: { __typename?: 'Space', id: any, name: string, description?: string | null } | null };

export type SpaceMembershipByPersonIdAndSpaceIdQueryVariables = Exact<{
  personId: Scalars['UUID'];
  spaceId: Scalars['UUID'];
}>;


export type SpaceMembershipByPersonIdAndSpaceIdQuery = { __typename?: 'Query', spaceMembershipByPersonIdAndSpaceId?: { __typename?: 'SpaceMembership', id: any } | null };

export type SpaceMembershipsByPersonIdQueryVariables = Exact<{
  personId: Scalars['UUID'];
}>;


export type SpaceMembershipsByPersonIdQuery = { __typename?: 'Query', spaceMemberships?: { __typename?: 'SpaceMembershipsConnection', edges: Array<{ __typename?: 'SpaceMembershipsEdge', node: { __typename?: 'SpaceMembership', id: any, space?: { __typename?: 'Space', id: any, name: string } | null } }> } | null };

export type SpaceMembershipsBySpaceIdQueryVariables = Exact<{
  spaceId: Scalars['UUID'];
}>;


export type SpaceMembershipsBySpaceIdQuery = { __typename?: 'Query', spaceMemberships?: { __typename?: 'SpaceMembershipsConnection', edges: Array<{ __typename?: 'SpaceMembershipsEdge', node: { __typename?: 'SpaceMembership', id: any, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string } | null } }> } | null };

export type SpacePostsQueryVariables = Exact<{
  spaceId: Scalars['UUID'];
  limit?: InputMaybe<Scalars['Int']>;
}>;


export type SpacePostsQuery = { __typename?: 'Query', posts?: { __typename?: 'PostsConnection', nodes: Array<{ __typename?: 'Post', id: any, body: string, membership?: { __typename?: 'SpaceMembership', id: any, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string } | null } | null }> } | null };

export type SpacePostsAddedSubscriptionVariables = Exact<{
  spaceId: Scalars['UUID'];
}>;


export type SpacePostsAddedSubscription = { __typename?: 'Subscription', posts?: { __typename?: 'PostSubscriptionPayload', event?: string | null, post?: { __typename?: 'Post', id: any, body: string, membership?: { __typename?: 'SpaceMembership', id: any, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string } | null } | null } | null } | null };

export type FamilyMembershipItemFragment = { __typename?: 'FamilyMembership', id: any, role: string, title?: string | null, person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string, user?: { __typename?: 'User', id: any } | null } | null } & { ' $fragmentName': 'FamilyMembershipItemFragment' };

export type CurrentPersonFamilyMembershipQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentPersonFamilyMembershipQuery = { __typename?: 'Query', currentPerson?: { __typename?: 'Person', id: any, familyMembership?: { __typename?: 'FamilyMembership', id: any, role: string, family?: { __typename?: 'Family', id: any, familyMemberships: { __typename?: 'FamilyMembershipsConnection', edges: Array<{ __typename?: 'FamilyMembershipsEdge', node: (
              { __typename?: 'FamilyMembership' }
              & { ' $fragmentRefs': { 'FamilyMembershipItemFragment': FamilyMembershipItemFragment } }
            ) }> } } | null } | null } | null };

export type CreateSpaceMembershipMutationVariables = Exact<{
  spaceId: Scalars['UUID'];
  personId: Scalars['UUID'];
}>;


export type CreateSpaceMembershipMutation = { __typename?: 'Mutation', createSpaceMembership?: { __typename?: 'CreateSpaceMembershipPayload', clientMutationId?: string | null } | null };

export type SharedSpacesQueryVariables = Exact<{
  person1: Scalars['UUID'];
  person2: Scalars['UUID'];
}>;


export type SharedSpacesQuery = { __typename?: 'Query', spaces?: { __typename?: 'SpacesConnection', edges: Array<{ __typename?: 'SpacesEdge', node: { __typename?: 'Space', id: any, name: string } }> } | null };

export type PersonPageDataQueryVariables = Exact<{
  id: Scalars['UUID'];
}>;


export type PersonPageDataQuery = { __typename?: 'Query', person?: { __typename?: 'Person', id: any, name: string, avatarUrl: string, createdAt: any } | null };

export const FamilyMembershipItemFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"FamilyMembershipItem"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"FamilyMembership"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<FamilyMembershipItemFragment, unknown>;
export const AllSpacesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"AllSpaces"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"personId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"spaces"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalCount"}},{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"spaceMemberships"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"condition"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"personId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"personId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<AllSpacesQuery, AllSpacesQueryVariables>;
export const CreateNewFamilyMemberDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateNewFamilyMember"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"role"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createNewFamilyMember"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"role"},"value":{"kind":"Variable","name":{"kind":"Name","value":"role"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"familyMembership"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"personId"}}]}}]}}]}}]} as unknown as DocumentNode<CreateNewFamilyMemberMutation, CreateNewFamilyMemberMutationVariables>;
export const CreateSpaceDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateSpace"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createSpace"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"space"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}]}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"space"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<CreateSpaceMutation, CreateSpaceMutationVariables>;
export const CurrentPersonDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentPerson"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentPerson"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}}]}}]}}]} as unknown as DocumentNode<CurrentPersonQuery, CurrentPersonQueryVariables>;
export const CurrentUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<CurrentUserQuery, CurrentUserQueryVariables>;
export const CurrentUserFamilyDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentUserFamily"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"family"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"familyMemberships"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orderBy"},"value":{"kind":"EnumValue","value":"CREATED_AT_ASC"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"nodes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"role"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<CurrentUserFamilyQuery, CurrentUserFamilyQueryVariables>;
export const CurrentFamilyMembershipDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentFamilyMembership"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentFamilyMembership"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"family"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]}}]} as unknown as DocumentNode<CurrentFamilyMembershipQuery, CurrentFamilyMembershipQueryVariables>;
export const CurrentUserWithManagedPeopleDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentUserWithManagedPeople"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentUser"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}}]}},{"kind":"Field","name":{"kind":"Name","value":"managedPeople"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"nodes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<CurrentUserWithManagedPeopleQuery, CurrentUserWithManagedPeopleQueryVariables>;
export const JoinSpaceDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"JoinSpace"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"personId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createSpaceMembership"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"spaceMembership"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"personId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"personId"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"spaceId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"roleId"},"value":{"kind":"StringValue","value":"member","block":false}}]}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"clientMutationId"}}]}}]}}]} as unknown as DocumentNode<JoinSpaceMutation, JoinSpaceMutationVariables>;
export const PostMessageDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"PostMessage"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"membershipId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"body"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"postMessage"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"spaceMembershipId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"membershipId"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"body"},"value":{"kind":"Variable","name":{"kind":"Name","value":"body"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"post"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<PostMessageMutation, PostMessageMutationVariables>;
export const SetPersonAvatarDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"SetPersonAvatar"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"personId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"avatarUrl"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updatePerson"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"personId"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"patch"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"avatarUrl"},"value":{"kind":"Variable","name":{"kind":"Name","value":"avatarUrl"}}}]}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"clientMutationId"}}]}}]}}]} as unknown as DocumentNode<SetPersonAvatarMutation, SetPersonAvatarMutationVariables>;
export const SpaceDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Space"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"space"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]} as unknown as DocumentNode<SpaceQuery, SpaceQueryVariables>;
export const SpaceMembershipByPersonIdAndSpaceIdDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SpaceMembershipByPersonIdAndSpaceId"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"personId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"spaceMembershipByPersonIdAndSpaceId"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"personId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"personId"}}},{"kind":"Argument","name":{"kind":"Name","value":"spaceId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<SpaceMembershipByPersonIdAndSpaceIdQuery, SpaceMembershipByPersonIdAndSpaceIdQueryVariables>;
export const SpaceMembershipsByPersonIdDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SpaceMembershipsByPersonId"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"personId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"spaceMemberships"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"condition"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"personId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"personId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"space"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<SpaceMembershipsByPersonIdQuery, SpaceMembershipsByPersonIdQueryVariables>;
export const SpaceMembershipsBySpaceIdDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SpaceMembershipsBySpaceId"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"spaceMemberships"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"condition"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"spaceId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<SpaceMembershipsBySpaceIdQuery, SpaceMembershipsBySpaceIdQueryVariables>;
export const SpacePostsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SpacePosts"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"limit"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"posts"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"condition"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"spaceId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}}}]}},{"kind":"Argument","name":{"kind":"Name","value":"last"},"value":{"kind":"Variable","name":{"kind":"Name","value":"limit"}}},{"kind":"Argument","name":{"kind":"Name","value":"orderBy"},"value":{"kind":"EnumValue","value":"CREATED_AT_ASC"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"nodes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"membership"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<SpacePostsQuery, SpacePostsQueryVariables>;
export const SpacePostsAddedDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"subscription","name":{"kind":"Name","value":"SpacePostsAdded"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"posts"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"spaceId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"event"}},{"kind":"Field","name":{"kind":"Name","value":"post"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"body"}},{"kind":"Field","name":{"kind":"Name","value":"membership"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"person"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<SpacePostsAddedSubscription, SpacePostsAddedSubscriptionVariables>;
export const CurrentPersonFamilyMembershipDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"CurrentPersonFamilyMembership"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currentPerson"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"familyMembership"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"role"}},{"kind":"Field","name":{"kind":"Name","value":"family"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"familyMemberships"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orderBy"},"value":{"kind":"EnumValue","value":"CREATED_AT_ASC"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"FamilyMembershipItem"}}]}}]}}]}}]}}]}}]}}]}},...FamilyMembershipItemFragmentDoc.definitions]} as unknown as DocumentNode<CurrentPersonFamilyMembershipQuery, CurrentPersonFamilyMembershipQueryVariables>;
export const CreateSpaceMembershipDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateSpaceMembership"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"personId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createSpaceMembership"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"spaceMembership"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"personId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"personId"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"spaceId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"spaceId"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"roleId"},"value":{"kind":"StringValue","value":"member","block":false}}]}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"clientMutationId"}}]}}]}}]} as unknown as DocumentNode<CreateSpaceMembershipMutation, CreateSpaceMembershipMutationVariables>;
export const SharedSpacesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SharedSpaces"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"person1"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"person2"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"spaces"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"filter"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"and"},"value":{"kind":"ListValue","values":[{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"spaceMemberships"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"some"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"personId"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"equalTo"},"value":{"kind":"Variable","name":{"kind":"Name","value":"person1"}}}]}}]}}]}}]},{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"spaceMemberships"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"some"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"personId"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"equalTo"},"value":{"kind":"Variable","name":{"kind":"Name","value":"person2"}}}]}}]}}]}}]}]}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"edges"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"node"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]} as unknown as DocumentNode<SharedSpacesQuery, SharedSpacesQueryVariables>;
export const PersonPageDataDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"PersonPageData"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"person"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"avatarUrl"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<PersonPageDataQuery, PersonPageDataQueryVariables>;