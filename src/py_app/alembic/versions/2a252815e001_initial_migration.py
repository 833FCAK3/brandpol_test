"""Initial migration

Revision ID: 2a252815e001
Revises: 
Create Date: 2024-01-31 13:08:13.197170

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = '2a252815e001'
down_revision: Union[str, None] = None
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_table('greetings',
    sa.Column('id', sa.Integer(), nullable=False),
    sa.Column('name', sa.String(), nullable=True),
    sa.Column('date', sa.DateTime(), nullable=True),
    sa.PrimaryKeyConstraint('id')
    )
    op.create_index(op.f('ix_greetings_id'), 'greetings', ['id'], unique=False)
    op.create_index(op.f('ix_greetings_name'), 'greetings', ['name'], unique=False)
    # ### end Alembic commands ###


def downgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_index(op.f('ix_greetings_name'), table_name='greetings')
    op.drop_index(op.f('ix_greetings_id'), table_name='greetings')
    op.drop_table('greetings')
    # ### end Alembic commands ###